<script lang="ts">
  import { t } from 'svelte-i18n';

  import { chainConfig } from '$chainConfig';
  import { Alert } from '$components/Alert';
  import ActionButton from '$components/Button/ActionButton.svelte';
  import { Card } from '$components/Card';
  import { ChainSelector, ChainSelectorDirection, ChainSelectorType } from '$components/ChainSelectors';
  import { errorToast, successToast, warningToast } from '$components/NotificationToast';
  import { infoToast } from '$components/NotificationToast/NotificationToast.svelte';
  import { TokenDropdown } from '$components/TokenDropdown';
  import { type Token, TokenType } from '$libs/token';
  import { connectedSourceChain } from '$stores';
  import { account } from '$stores';

  // Extended Token interface for faucet
  interface FaucetToken extends Token {
    amount: string;
    faucetType: 'ERC20' | 'ERC721' | 'NATIVE';
    tokenAddress: string;
  }

  // Get token list from API

  let claiming = false;
  let loadingTokens = false;

  let selectedToken: FaucetToken | undefined;
  let claimButtonEnabled = true;
  let alertMessage = '';
  let availableTokens: FaucetToken[] = [];
  let lastFetchedChainId: number | null = null;

  $: currentChainId = $connectedSourceChain?.id;
  //   $: currentChain = chains.find((chain) => chain.id === currentChainId);
  //   $: currencySymbol = currentChain?.nativeCurrency.symbol || '';
  $: connected = isUserConnected($account);
  $: disabled = !$account || !$account.isConnected;

  // Backend API call function
  async function claimTokenFromApi(tokenConfig: FaucetToken, userAddress: string, chainId: number) {
    const response = await fetch('/api/faucet/claim', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        tokenAddress: tokenConfig.tokenAddress,
        tokenSymbol: tokenConfig.symbol,
        amount: tokenConfig.amount,
        decimals: tokenConfig.decimals,
        logoUri: tokenConfig.logoURI,
        userAddress,
        chainId,
        tokenType: tokenConfig.faucetType,
      }),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to claim token');
    }

    return await response.json();
  }

  async function claimToken() {
    if (claiming) return;
    if (!selectedToken || !$connectedSourceChain || !$account?.address) return;

    claiming = true;

    try {
      const result = await claimTokenFromApi(selectedToken, $account.address, $connectedSourceChain.id);

      const explorer = chainConfig[$connectedSourceChain.id]?.blockExplorers?.default.url;

      infoToast({
        title: $t('faucet.mint.tx.title'),
        message: $t('faucet.mint.tx.message', {
          values: {
            token: selectedToken.symbol,
            url: `${explorer}/tx/${result.txHash}`,
          },
        }),
      });

      successToast({
        title: $t('faucet.mint.success.title'),
        message: $t('faucet.mint.success.message'),
      });
    } catch (err) {
      console.error(err);
      const error = err as Error;

      if (error.message?.includes('already claimed')) {
        warningToast({
          title: 'Already Claimed',
          message: 'You have already claimed this token, please wait for cooldown period to end',
        });
      } else if (error.message?.includes('insufficient')) {
        errorToast({
          title: 'Insufficient Balance',
          message: 'Faucet contract has insufficient balance, please contact administrator',
        });
      } else {
        errorToast({
          title: 'Claim Failed',
          message: error.message || 'Unknown error',
        });
      }
    } finally {
      claiming = false;
    }
  }

  function isUserConnected(user: Maybe<typeof $account>) {
    return Boolean(user?.isConnected);
  }

  const handleTokenSelected = (event: CustomEvent<{ token: Token }>) => {
    selectedToken = event.detail.token as FaucetToken;
  };

  function getAlertMessage(connected: boolean) {
    if (!connected) return $t('messages.account.required');
    return '';
  }

  // API call to get token list
  async function fetchTokens(chainId: number): Promise<FaucetToken[]> {
    if (!chainId) return [];

    loadingTokens = true;
    try {
      const response = await fetch(`/api/faucet/tokens?chainId=${chainId}`);
      if (!response.ok) {
        throw new Error('Failed to fetch token list');
      }
      const data = await response.json();
      // Convert API response to FaucetToken format
      return (data.tokens || []).map((token: FaucetToken) => ({
        symbol: token.symbol,
        name: token.name,
        decimals: token.decimals,
        addresses: token.addresses || {},
        logoURI: token.logoURI,
        type:
          token.faucetType === 'NATIVE'
            ? TokenType.ETH
            : token.faucetType === 'ERC20'
              ? TokenType.ERC20
              : TokenType.ERC721,
        // Add custom properties for faucet
        amount: token.amount,
        faucetType: token.faucetType,
        tokenAddress: token.tokenAddress,
      }));
    } catch (error) {
      console.error('Failed to fetch token list:', error);
      return [];
    } finally {
      loadingTokens = false;
    }
  }

  // Listen to network changes, update available tokens
  $: if (currentChainId && currentChainId !== lastFetchedChainId) {
    lastFetchedChainId = currentChainId;
    fetchTokens(currentChainId).then((tokens) => {
      availableTokens = tokens;

      // Handle selected token update logic
      if (selectedToken) {
        // Find token with same symbol in new network
        const matchingToken = tokens.find((t: FaucetToken) => t.symbol === selectedToken?.symbol);

        if (matchingToken) {
          // If found, update to new network configuration
          selectedToken = matchingToken;
        } else {
          // If not found in new network, clear selection
          selectedToken = undefined;
        }
      }
    });
  } else if (!currentChainId) {
    availableTokens = [];
    selectedToken = undefined;
    lastFetchedChainId = null;
  }

  $: alertMessage = getAlertMessage(connected);
</script>

<Card class="w-full md:w-[524px]" title="Faucet" text="Mint test tokens.">
  <div class="space-y-[35px]">
    <div class="space-y-2">
      <ChainSelector
        type={ChainSelectorType.SMALL}
        direction={ChainSelectorDirection.SOURCE}
        label={$t('chain_selector.currently_on')}
        switchWallet />

      {#if loadingTokens}
        <div class="text-center py-4">
          <span class="text-sm text-gray-600">Loading token list...</span>
        </div>
      {:else if availableTokens.length > 0}
        <TokenDropdown
          {disabled}
          tokens={availableTokens}
          onlyMintable={true}
          bind:value={selectedToken}
          on:tokenSelected={handleTokenSelected} />
      {:else if currentChainId}
        <Alert type="warning">
          Current network (Chain ID: {currentChainId}) does not support token claiming
        </Alert>
      {/if}
    </div>

    {#if alertMessage}
      <Alert type="warning" forceColumnFlow>
        {alertMessage}
      </Alert>
    {/if}

    {#if availableTokens.length > 0}
      <ActionButton
        priority="primary"
        disabled={!claimButtonEnabled || disabled || !selectedToken}
        loading={claiming}
        on:click={claimToken}>
        <span class="body-bold">
          {#if claiming}
            Claiming...
          {:else}
            Claim {selectedToken ? `${selectedToken.amount} ${selectedToken.symbol}` : 'Token'}
          {/if}
        </span>
      </ActionButton>
    {/if}

    {#if selectedToken}
      <div class="text-sm text-gray-600">
        <p>Token Type: {selectedToken.faucetType === 'NATIVE' ? 'Native Token' : 'ERC20 Token'}</p>
        {#if selectedToken.faucetType !== 'NATIVE'}
          <p>Token Address: {selectedToken.tokenAddress}</p>
        {/if}
        <p>Claim Amount: {selectedToken.amount} {selectedToken.symbol}</p>
      </div>
    {/if}
  </div>
</Card>
