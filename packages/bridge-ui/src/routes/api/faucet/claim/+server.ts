import { json } from '@sveltejs/kit';
import { type Chain, createPublicClient, createWalletClient, http, parseEther, parseUnits } from 'viem';
import { privateKeyToAccount } from 'viem/accounts';
import { arbitrumSepolia } from 'viem/chains';

import { env } from '$env/dynamic/private';
import { getLogger } from '$libs/util/logger';

import { _TOKEN_CONFIG } from '../tokens/+server.js';
import type { RequestHandler } from './$types';

const log = getLogger('faucet:claim');

// Use Upstash Redis as serverless storage
// Create a simple implementation without installing @upstash/redis
// For production, please install: npm install @upstash/redis

// Redis configuration - read from environment variables
const REDIS_CONFIG = {
  url: env.KV_REST_API_URL || '',
  token: env.KV_REST_API_TOKEN || '',
};

// Define Moonchain Geneva testnet
const moonchainGeneva = {
  id: 5167004,
  name: 'Moonchain Geneva Testnet',
  network: 'moonchain-geneva',
  nativeCurrency: {
    name: 'MXC',
    symbol: 'MXC',
    decimals: 18,
  },
  rpcUrls: {
    default: {
      http: ['https://geneva-rpc.moonchain.com'],
    },
    public: {
      http: ['https://geneva-rpc.moonchain.com'],
    },
  },
  blockExplorers: {
    default: {
      name: 'Moonchain Geneva Explorer',
      url: 'https://geneva-explorer.moonchain.com',
    },
  },
  testnet: true,
} as const;

// Network configuration
const NETWORK_CONFIG: Record<
  number,
  {
    chain: Chain;
    rpcUrl: string;
    privateKeyEnvVar: string;
  }
> = {
  421614: {
    chain: arbitrumSepolia,
    rpcUrl: 'https://sepolia-rollup.arbitrum.io/rpc',
    privateKeyEnvVar: env.FAUCET_PRIVATE_KEY || '',
  },
  5167004: {
    chain: moonchainGeneva,
    rpcUrl: 'https://geneva-rpc.moonchain.com',
    privateKeyEnvVar: env.FAUCET_PRIVATE_KEY || '',
  },
};

// ERC20 ABI
const ERC20_ABI = [
  {
    name: 'transfer',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'to', type: 'address' },
      { name: 'amount', type: 'uint256' },
    ],
    outputs: [{ name: '', type: 'bool' }],
  },
  {
    name: 'balanceOf',
    type: 'function',
    stateMutability: 'view',
    inputs: [{ name: 'account', type: 'address' }],
    outputs: [{ name: '', type: 'uint256' }],
  },
] as const;

// Cooldown period (24 hours in seconds)
const COOLDOWN_PERIOD_SECONDS = 24 * 60 * 60;

// Maximum claims per day
const MAX_CLAIMS_PER_DAY = 1;

// Send lock time (5 minutes, prevent duplicate sending)
const SEND_LOCK_SECONDS = 5 * 60;

// Simple Redis REST API client
class SimpleRedisClient {
  private url: string;
  private token: string;

  constructor(url: string, token: string) {
    this.url = url;
    this.token = token;
  }

  private async request(command: string[]) {
    if (!this.url || !this.token) {
      throw new Error('Redis Error: Redis configuration is missing');
    }

    const response = await fetch(`${this.url}`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${this.token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(command),
    });

    if (!response.ok) {
      throw new Error(`Redis request failed: ${response.statusText}`);
    }

    const data = await response.json();
    return data.result;
  }

  async get(key: string): Promise<string | null> {
    return await this.request(['GET', key]);
  }

  async set(key: string, value: string, exSeconds?: number): Promise<void> {
    if (exSeconds) {
      await this.request(['SET', key, value, 'EX', exSeconds.toString()]);
    } else {
      await this.request(['SET', key, value]);
    }
  }

  async incr(key: string): Promise<number> {
    return await this.request(['INCR', key]);
  }

  async expire(key: string, seconds: number): Promise<void> {
    await this.request(['EXPIRE', key, seconds.toString()]);
  }

  async ttl(key: string): Promise<number> {
    return await this.request(['TTL', key]);
  }

  async del(key: string): Promise<void> {
    await this.request(['DEL', key]);
  }
}

const redis = new SimpleRedisClient(REDIS_CONFIG.url, REDIS_CONFIG.token);

function getClaimKey(userAddress: string, tokenAddress: string, chainId: number) {
  return `faucet:${userAddress.toLowerCase()}:${tokenAddress.trim().toLowerCase()}:${chainId}`;
}

function getDailyCountKey(userAddress: string, tokenAddress: string, chainId: number) {
  const today = new Date().toISOString().split('T')[0]; // YYYY-MM-DD
  return `faucet:daily:${today}:${userAddress.toLowerCase()}:${tokenAddress.trim().toLowerCase()}:${chainId}`;
}

function getSendLockKey(userAddress: string, tokenAddress: string, chainId: number) {
  return `faucet:sending:${userAddress.toLowerCase()}:${tokenAddress.trim().toLowerCase()}:${chainId}`;
}

async function canClaim(userAddress: string, tokenAddress: string, chainId: number): Promise<boolean> {
  try {
    const claimKey = getClaimKey(userAddress, tokenAddress, chainId);
    const dailyCountKey = getDailyCountKey(userAddress, tokenAddress, chainId);
    const sendLockKey = getSendLockKey(userAddress, tokenAddress, chainId);

    // Check if currently sending
    const sendLock = await redis.get(sendLockKey);
    if (sendLock) {
      return false; // Currently sending, prevent duplicate claims
    }

    // Check cooldown period
    const lastClaimTime = await redis.get(claimKey);
    if (lastClaimTime) {
      const ttl = await redis.ttl(claimKey);
      if (ttl > 0) {
        return false; // Still in cooldown period
      }
    }

    // Check daily limit
    const dailyCount = await redis.get(dailyCountKey);
    if (dailyCount && parseInt(dailyCount) >= MAX_CLAIMS_PER_DAY) {
      return false; // Exceeded daily limit
    }

    return true;
  } catch (error) {
    log('Failed to check claim status: %o', error);
    // If Redis fails, allow claiming for better user experience but log the error
    return true;
  }
}

// Set send lock
async function setSendLock(userAddress: string, tokenAddress: string, chainId: number): Promise<void> {
  try {
    const sendLockKey = getSendLockKey(userAddress, tokenAddress, chainId);
    await redis.set(sendLockKey, 'sending', SEND_LOCK_SECONDS);
  } catch (error) {
    log('Failed to set send lock: %o', error);
  }
}

// Release send lock
async function releaseSendLock(userAddress: string, tokenAddress: string, chainId: number): Promise<void> {
  try {
    const sendLockKey = getSendLockKey(userAddress, tokenAddress, chainId);
    await redis.del(sendLockKey);
  } catch (error) {
    log('Failed to release send lock: %o', error);
  }
}

async function updateClaimRecord(userAddress: string, tokenAddress: string, chainId: number): Promise<void> {
  try {
    const claimKey = getClaimKey(userAddress, tokenAddress, chainId);
    const dailyCountKey = getDailyCountKey(userAddress, tokenAddress, chainId);
    const now = Date.now().toString();

    // Set cooldown period
    await redis.set(claimKey, now, COOLDOWN_PERIOD_SECONDS);

    // Increase daily count, set to expire at midnight tomorrow
    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);
    tomorrow.setHours(0, 0, 0, 0);
    const secondsUntilMidnight = Math.floor((tomorrow.getTime() - Date.now()) / 1000);

    const currentCount = await redis.incr(dailyCountKey);
    if (currentCount === 1) {
      // First count, set expiration time
      await redis.expire(dailyCountKey, secondsUntilMidnight);
    }
  } catch (error) {
    log('Failed to update claim record: %o', error);
    // Don't block transaction even if record update fails
  }
}

// Send tokens to blockchain (supports native token and ERC20)
async function sendToken(
  tokenAddress: string,
  tokenType: string,
  userAddress: string,
  amount: string,
  decimals: number,
  chainId: number,
): Promise<string> {
  // Validate token configuration
  const chainTokens = _TOKEN_CONFIG[chainId];
  if (!chainTokens) {
    throw new Error(`Unsupported network: ${chainId}`);
  }

  const tokenConfig = chainTokens.find((t) => t.tokenAddress === tokenAddress && t.faucetType === tokenType);

  if (!tokenConfig) {
    throw new Error('Token configuration not found');
  }

  // Validate requested amount matches configuration
  if (tokenConfig.amount !== amount) {
    throw new Error(`Amount mismatch, should be ${tokenConfig.amount}`);
  }

  // Get network configuration
  const networkConfig = NETWORK_CONFIG[chainId];
  if (!networkConfig) {
    throw new Error(`Unsupported network: ${chainId}`);
  }

  // Get private key
  const privateKey = networkConfig.privateKeyEnvVar;
  if (!privateKey) {
    throw new Error(`Private key not configured: ${networkConfig.privateKeyEnvVar}`);
  }

  // Create account and clients
  const account = privateKeyToAccount(`0x${privateKey}` as `0x${string}`);

  const publicClient = createPublicClient({
    chain: networkConfig.chain,
    transport: http(networkConfig.rpcUrl),
  });

  const walletClient = createWalletClient({
    account,
    chain: networkConfig.chain,
    transport: http(networkConfig.rpcUrl),
  });

  let hash: string;

  try {
    if (tokenType === 'NATIVE') {
      // Send native token (ETH/MXC etc.)
      const value = parseEther(amount);

      hash = await walletClient.sendTransaction({
        to: userAddress as `0x${string}`,
        value,
        chain: null,
      });

      log('Sent %s native %s to %s (chain %d), hash: %s', amount, tokenConfig.symbol, userAddress, chainId, hash);
    } else if (tokenType === 'ERC20') {
      // Send ERC20 token
      const transferAmount = parseUnits(amount, decimals);

      // Check faucet wallet balance
      const balance = await publicClient.readContract({
        address: tokenAddress as `0x${string}`,
        abi: ERC20_ABI,
        functionName: 'balanceOf',
        args: [account.address],
      });

      if (balance < transferAmount) {
        throw new Error('insufficient');
      }

      // Send ERC20 token
      hash = await walletClient.writeContract({
        address: tokenAddress as `0x${string}`,
        abi: ERC20_ABI,
        functionName: 'transfer',
        args: [userAddress as `0x${string}`, transferAmount],
        chain: null,
      });

      log('Sent %s %s to %s (chain %d), hash: %s', amount, tokenConfig.symbol, userAddress, chainId, hash);
    } else {
      throw new Error(`Unsupported token type: ${tokenType}`);
    }

    // Wait for transaction confirmation
    // await publicClient.waitForTransactionReceipt({ hash: hash as `0x${string}` });

    return hash;
  } catch (error) {
    log('Failed to send token: %o', error);
    throw error;
  }
}

// POST request - Claim tokens
export const POST: RequestHandler = async ({ request }) => {
  try {
    const { tokenAddress, tokenSymbol, amount, decimals, userAddress, chainId, tokenType } = await request.json();

    // Validate parameters
    if (!tokenAddress || !userAddress || !amount || !chainId || decimals === undefined) {
      return json({ error: 'Missing required parameters' }, { status: 400 });
    }

    // Validate address format (simple check)
    if ((!tokenAddress.startsWith('0x') && tokenType !== 'NATIVE') || !userAddress.startsWith('0x')) {
      log(
        'Invalid address format - tokenAddress: %s, tokenType: %s, userAddress: %s',
        tokenAddress,
        tokenType,
        userAddress,
      );
      return json({ error: 'Invalid address format' }, { status: 400 });
    }

    // Check if can claim
    const claimAllowed = await canClaim(userAddress, tokenAddress, chainId);
    if (!claimAllowed) {
      return json({ error: 'already claimed' }, { status: 429 });
    }

    // Set send lock
    await setSendLock(userAddress, tokenAddress, chainId);

    try {
      // Send token
      const txHash = await sendToken(tokenAddress, tokenType, userAddress, amount, decimals, chainId);

      // Update claim record
      await updateClaimRecord(userAddress, tokenAddress, chainId);

      return json({
        success: true,
        txHash,
        amount,
        symbol: tokenSymbol,
        message: `Successfully sent ${amount} ${tokenSymbol} to ${userAddress}`,
      });
    } catch (sendError) {
      log('Failed to send token: %o', sendError);
      const sendErrorMsg = (sendError as Error).message || 'Unknown error';

      if (sendErrorMsg.includes('insufficient')) {
        return json({ error: 'insufficient' }, { status: 503 });
      }

      return json({ error: sendErrorMsg }, { status: 500 });
    } finally {
      // Release send lock
      await releaseSendLock(userAddress, tokenAddress, chainId);
    }
  } catch (error) {
    log('Failed to claim token: %o', error);
    const errorMsg = (error as Error).message || 'Unknown error';

    return json({ error: errorMsg }, { status: 500 });
  }
};
