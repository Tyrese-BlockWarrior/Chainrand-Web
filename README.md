# Chainrand-web

This repo contains the implementation of Chainrand.

# Smart Contract 

Contains functionality to tie a Chainlink VRF to off-chain code in an immutable manner.

- Functions to mint and update the Chainrand NFT.

- Methods for bulk querying of token data to minimize RPC calls.

- Efficient on-chain SVG art for NFTs without a custom image.

**Files**

- `contract`

# Frontend

Single-Page-App based frontend. 

- Just upload and serve from any webserver. 

**Files**

- `css`
- `js`
- `assets`
- `fonts`
- `index.html`
- `favicon.ico`
	
# Backend

Golang based backend for quick querying of the Chainrand NFT based on token name or minter. 

The search function serves several purposes:

- Discouraging spam

	- At Chainrand, we do not have a paid name reservation system to ensure maximum affordability. However, this invites potential spam. We do not want a situation where search results are flooded with fake collections.

	- A solution is to provide a search based on a veracity score. 
		
		- If two tokens share the exact same name, the one that is verified will be ranked higher. 
		
		- If both tokens are not verified, then the one that paid a higher nominal fee will be ranked higher.

- Future monetization avenue

	- The search ranking system can also provide monetization in the future to cover server fees.

- Cherry-picking detection

	- If there are many tokens with the similar names that did not pay a decent nominal fee, it is likely that the minter(s) are trying to do cherry-picking.

**Files**

- `api`

# License

MIT

