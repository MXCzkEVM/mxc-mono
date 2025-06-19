import { json } from '@sveltejs/kit';

import { env } from '$env/dynamic/private';
import { getLogger } from '$libs/util/logger';

import type { RequestHandler } from './$types';

const log = getLogger('faucet:tokens');

// Default token configuration - fallback if environment variable is not set
const DEFAULT_TOKEN_CONFIG: Record<
  number,
  Array<{
    symbol: string;
    name: string;
    tokenAddress: string; // Use 'native' for native tokens
    amount: string; // Claim amount
    decimals: number;
    faucetType: 'ERC20' | 'ERC721' | 'NATIVE';
    logoURI?: string; // Optional icon URL, will be auto-generated if not provided
  }>
> = {};

// Load token configuration from environment variable or use default
function getTokenConfig(): Record<
  number,
  Array<{
    symbol: string;
    name: string;
    tokenAddress: string;
    amount: string;
    decimals: number;
    faucetType: 'ERC20' | 'ERC721' | 'NATIVE';
    logoURI?: string;
  }>
> {
  try {
    const configJson = env.FAUCET_TOKEN_CONFIG;
    if (configJson) {
      const parsed = JSON.parse(configJson);
      log('Loaded token config from environment variable');
      return parsed;
    }
  } catch (error) {
    log('Failed to parse FAUCET_TOKEN_CONFIG environment variable: %o', error);
    log('Using default token configuration');
  }

  return DEFAULT_TOKEN_CONFIG;
}

// Token configuration - loaded from environment variable or default
export const _TOKEN_CONFIG = getTokenConfig();

// GET request - Get claimable token list for specified network
export const GET: RequestHandler = async ({ url }) => {
  try {
    const chainId = parseInt(url.searchParams.get('chainId') || '0');

    if (!chainId) {
      return json({ error: 'Please provide valid chainId parameter' }, { status: 400 });
    }

    const tokens = _TOKEN_CONFIG[chainId] || [];

    return json({
      success: true,
      chainId,
      tokens: tokens.map((token) => ({
        ...token,
        // Format address for frontend components
        addresses: token.tokenAddress === 'native' ? {} : { [chainId]: token.tokenAddress },
      })),
    });
  } catch (error) {
    log('Failed to get token list: %o', error);
    return json({ error: 'Failed to get token list' }, { status: 500 });
  }
};
