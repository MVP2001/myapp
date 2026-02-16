#!/bin/bash
# Initialize Vault for production
set -e

echo "🔐 Initializing Vault..."

until docker-compose exec -T vault vault status > /dev/null 2>&1; do
    echo "Waiting for Vault to start..."
    sleep 2
done

INIT_OUTPUT=$(docker-compose exec -T vault vault operator init -key-shares=3 -key-threshold=2 -format=json)

echo "$INIT_OUTPUT" | jq -r '.root_token' > vault/root-token.txt
echo "$INIT_OUTPUT" | jq -r '.unseal_keys_b64[]' > vault/unseal-keys.txt

echo "✅ Vault initialized!"
echo "⚠️  IMPORTANT: Store unseal keys and root token securely!"

