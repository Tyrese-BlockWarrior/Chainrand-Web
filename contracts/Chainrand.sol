// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
pragma abicoder v2;

import "@chainlink/contracts@0.2.2/src/v0.8/interfaces/LinkTokenInterface.sol";
import "@chainlink/contracts@0.2.2/src/v0.8/VRFRequestIDBase.sol";
import "@chainlink/contracts@0.2.2/src/v0.8/VRFConsumerBase.sol";

import "@openzeppelin/contracts@4.3.2/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts@4.3.2/access/Ownable.sol";
import "@openzeppelin/contracts@4.3.2/security/ReentrancyGuard.sol";

import "base64-sol/base64.sol";

contract Chainrand is VRFConsumerBase, ERC721Enumerable, Ownable, ReentrancyGuard {
    
    using Strings for uint;

    /// @dev This is the max uint that can be safely represented in JS double.
    uint public MAX_TOKENS = 9007199254740991;

    mapping(bytes32 => uint) internal vrfs;

    /// @dev Metadata struct for token 
    struct Token {
        string name; // fixed upon mint
        string codeURI; // fixed upon mint
        string seedKey; 
        string imageURI;
        string projectURI;  
        address minter; // fixed upon mint
        uint seedKeyHash; // fixed upon mint
        uint codeHash; // fixed upon mint
        uint randomness; // fixed upon mint (upon VRF fufilled)
        uint paid; 
        bool verified;
    }

    /// @dev Metadata of tokens
    mapping(uint => Token) public tokens;

    /// @dev Nominal minimum mint fee. 
    uint public mintFee;

    bytes32 internal vrfkeyHash;

    uint public vrfFee;

    event Verified(uint _tokenId);

    event VerificationRevoked(uint _tokenId);

    event RandomnessFufilled(uint _tokenId);

    event SeedKeyRevealed(uint _tokenId);

    event Paid(uint _tokenId);

    // Get the values to connect to Chainlink from
    // https://docs.chain.link/docs/vrf-contracts/
    //
    // Ethereum Mainnet
    // ----------------
    // LINK Token: 0x514910771AF9Ca656af840dff83E8264EcF986CA
    // VRF Coordinator: 0xf0d54349aDdcf704F77AE15b96510dEA15cb7952
    // Key Hash: 0xAA77729D3466CA35AE8D28B3BBAC7CC36A5031EFDC430821C02BC31A238AF445 
    // Fee: 2
    //
    // Rinkeby Testnet
    // ---------------
    // LINK Token: 0x01BE23585060835E02B77ef475b0Cc51aA1e0709
    // VRF Coordinator: 0xb3dCcb4Cf7a26f6cf6B120Cf5A73875B7BBc655B
    // Key Hash: 0x2ed0feb3e7fd2022120aa84fab1945545a9f2ffc9076fd6156fa96eaff4c1311
    // Fee: 0.1
    // 
    // Polygon Mainnet
    // ---------------
    // LINK Token: 0xb0897686c545045aFc77CF20eC7A532E3120E0F1
    // VRF Coordinator: 0x3d2341ADb2D31f1c5530cDC622016af293177AE0
    // Key Hash: 0xf86195cf7690c55907b2b611ebb7343a6f649bff128701cc542f0569e2c549da
    // Fee: 0.0001
    //
    // Mumbai Testnet
    // --------------
    // LINK Token: 0x326C977E6efc84E512bB9C30f76E30c160eD06FB
    // VRF Coordinator: 0x8C7382F9D8f56b33781fE506E897a4F1e2d17255
    // Key Hash: 0x6e75b569a01ef56d18cab6a8e71e6600d6ce853834d4a5748b720d06f878b3a4
    // Fee: 0.0001
     
    constructor() 
    ERC721("Chainrand", "RAND") 
    VRFConsumerBase(
        0xb3dCcb4Cf7a26f6cf6B120Cf5A73875B7BBc655B, // VRF Coordinator
        0x01BE23585060835E02B77ef475b0Cc51aA1e0709  // LINK Token
        )
    { 
        vrfkeyHash = 0x2ed0feb3e7fd2022120aa84fab1945545a9f2ffc9076fd6156fa96eaff4c1311;
        vrfFee = 0.1 * 10 ** 18; // 0.1 LINK (Varies by network)
    }

    /// @dev Mint a Chainrand token.
    function mint(string memory _name, uint _seedKeyHash, uint _codeHash, 
        string memory _codeURI, string memory _imageURI, string memory _projectURI) 
    public payable nonReentrant {
        require(LINK.transferFrom(msg.sender, address(this), vrfFee), "LINK payment failed.");
        require(mintFee <= msg.value, "Insufficient payment.");
        
        uint tokenId = totalSupply();
        require(tokenId < MAX_TOKENS, "No more tokens available.");
        
        _safeMint(msg.sender, tokenId);
        tokens[tokenId] = Token(_name, _codeURI, "", _imageURI, _projectURI, 
            msg.sender, _seedKeyHash, _codeHash, 0, msg.value, false);
        vrfs[requestRandomness(vrfkeyHash, vrfFee)] = tokenId;
    }

    // Callback function used by VRF Coordinator
    function fulfillRandomness(bytes32 _requestId, uint _randomness) internal override {
        uint tokenId = vrfs[_requestId];
        tokens[tokenId].randomness = _randomness;
        emit RandomnessFufilled(tokenId);
    }

    /// @dev For payments to support Chainrand. 
    function pay(uint _tokenId) public payable {
        require(_tokenId < totalSupply(), "Token not found.");
        tokens[_tokenId].paid += msg.value;
        emit Paid(_tokenId);
    }

    /// @dev Sets the mint fee.
    function setMintFee(uint _mintFee) public onlyOwner {
        mintFee = _mintFee;
    }

    // Fake Mint.... for testing purposes.
    function fakeMint(uint _count, string memory _name, uint _seedKeyHash, uint _codeHash, 
        string memory _codeURI, string memory _imageURI, string memory _projectURI) 
    public onlyOwner {
        unchecked {
            for (uint i = 0; i < _count; i++) {
                uint tokenId = totalSupply();
                _safeMint(msg.sender, tokenId);
                uint randomness = uint(keccak256(abi.encodePacked(tokenId, msg.sender, blockhash(block.number - 1))));
                tokens[tokenId] = Token(_name, _codeURI, "", _imageURI, _projectURI, 
                    msg.sender, _seedKeyHash, _codeHash, randomness, 0, false);
                emit RandomnessFufilled(tokenId);
            }
        }
    }

    /// @dev Withdraws Ether for the owner.    
    function withdraw() public onlyOwner {
        uint256 amount = address(this).balance;
        payable(msg.sender).transfer(amount);
    }

    /// @dev Change the token name.
    function setSeedKey(uint _tokenId, string memory _seedKey) public {
        require(ownerOf(_tokenId) == msg.sender, "You do not own this token.");
        require(sha256(bytes(_seedKey)) == bytes32(tokens[_tokenId].seedKeyHash), 
            "Seed key hash does not match.");
        tokens[_tokenId].seedKey = _seedKey;
        emit SeedKeyRevealed(_tokenId);
    }

    /// @dev Change the token's image URI
    function setImageURI(uint _tokenId, string memory _imageURI) public {
        require(ownerOf(_tokenId) == msg.sender, "You do not own this token.");
        tokens[_tokenId].imageURI = _imageURI;
    }

    /// @dev Change the token's project URI
    function setProjectURI(uint _tokenId, string memory _projectURI) public {
        require(ownerOf(_tokenId) == msg.sender, "You do not own this token.");
        tokens[_tokenId].projectURI = _projectURI;
    }

    function escapeJsonString(string memory symbol) internal pure returns (string memory) {
        unchecked {
            bytes memory symbolBytes = bytes(symbol);
            uint escapesLen = 0;
            for (uint i = 0; i < symbolBytes.length; i++) {
                uint b = uint(uint8(symbolBytes[i]));
                if (b == 34 || b == 92 || b <= 31) { 
                    escapesLen += 5;
                }
            }
            bytes16 alphabet = "0123456789abcdef";
            if (escapesLen > 0) {
                bytes memory escapedBytes = new bytes(symbolBytes.length + escapesLen);
                uint index;
                for (uint i = 0; i < symbolBytes.length; i++) {
                    uint b = uint(uint8(symbolBytes[i]));
                    if (b == 34 || b == 92 || b <= 31) {
                        escapedBytes[index++] = '\\';
                        escapedBytes[index++] = 'u';
                        escapedBytes[index++] = '0';
                        escapedBytes[index++] = '0';
                        escapedBytes[index++] = alphabet[b >> 4];
                        escapedBytes[index++] = alphabet[b & 15];
                    } else {
                        escapedBytes[index++] = symbolBytes[i];    
                    }
                }
                return string(escapedBytes);
            }
            return symbol;    
        }
    }

    function escapeCodeQuotes(string memory symbol) internal pure returns (string memory) {
        unchecked {
            bytes memory symbolBytes = bytes(symbol);
            uint escapesLen = 0;
            for (uint i = 0; i < symbolBytes.length; i++) {
                uint b = uint(uint8(symbolBytes[i]));
                if (b == 96) { 
                    escapesLen += 1;
                }
            }
            bytes memory escapedBytes = new bytes(symbolBytes.length + escapesLen);
            uint index;
            for (uint i = 0; i < symbolBytes.length; i++) {
                uint b = uint(uint8(symbolBytes[i]));
                if (b == 96) {
                    escapedBytes[index++] = '\\';
                    escapedBytes[index++] = '`';
                } else {
                    escapedBytes[index++] = symbolBytes[i];    
                }
            }
            return string(escapedBytes);
        }
    }

    function concat(string memory s0, string memory s1, string memory s2, string memory s3)
    internal pure returns (string memory) {
        unchecked { return string(abi.encodePacked(s0, s1, s2, s3)); }
    }
    
    function codeQuotes(string memory symbol) internal pure returns (string memory) {
        unchecked { return concat('`', symbol, '`', ''); }
    }
    
    function addressToString(address _addr) internal pure returns(string memory) {
        unchecked {
            bytes32 value = bytes32(uint(uint160(_addr)));
            bytes16 alphabet = "0123456789abcdef";
            bytes memory str = new bytes(42);
            str[0] = "0";
            str[1] = "x";
            for (uint i = 0; i < 20; i++) {
                uint b = uint(uint8(value[i + 12]));
                uint j = 2+i*2;
                str[j  ] = alphabet[b >> 4];
                str[j+1] = alphabet[b & 15];
            }
            return string(str);
        }
    }

    /// @dev Returns the token URI in base64 data json.
    function tokenURI(uint _tokenId) override public view returns (string memory) {
        unchecked {
            require(_tokenId < totalSupply(), "Token not found.");
            Token memory t = tokens[_tokenId];
            uint r = t.randomness;
            string memory image;
            if (bytes(t.imageURI).length == 0) {
                for (uint i = 0; i < 32; i++) {
                    uint d = i & 7;
                    string memory x = 
                    (d == 0) ? "<path d='M" :
                    (d == 2) ? "L" :
                    (d <= 3) ? " " :
                    (d == 4) ? "' style='stroke:rgba(" :
                    (d <= 6) ? "," : ",0.5);stroke-width:";
                    image = concat(image, x, (r & 255).toString(), (d == 7 ? "px'/>" : ""));
                    r >>= 8;
                }
                image = Base64.encode(abi.encodePacked(
                    "<svg xmlns='http://www.w3.org/2000/svg' version='1.1' viewBox='0 0 256 256'>", 
                    image, "</svg>"));
                image = concat("data:image/svg+xml;base64,", image, '', '');    
            } else {
                image = t.imageURI;
            }
            
            string memory a = concat(
                '\\n\\nVerified: ', codeQuotes(t.verified ? "Yes" : "No"), 
                '\\n\\nMinter: ', codeQuotes(addressToString(t.minter)) 
            );
            if (bytes(t.projectURI).length > 0) {
                a = concat(a, '\\n\\nProject URI: [Link](', escapeJsonString(t.projectURI),')');
            }
            if (bytes(t.codeURI).length > 0) {
                a = concat(a, '\\n\\nCode URI: [Link](', escapeJsonString(t.codeURI),')');
            }
            a = concat(a, '\\n\\nCode Hash: ', codeQuotes(t.codeHash.toHexString()), '');
            a = concat(a, '\\n\\nSeed Key Hash: ', codeQuotes(t.seedKeyHash.toHexString()), '');
            if (bytes(t.seedKey).length > 0) {
                a = concat(a, '\\n\\nSeed Key: ', codeQuotes(escapeCodeQuotes(t.seedKey)), '');
            }
            if (t.randomness > 0) {
                a = concat(a, '\\n\\nRandomness: ', codeQuotes(t.randomness.toString()), '');
            }

            string memory json = Base64.encode(abi.encodePacked(
                '{"name":"', t.name, '","description":"Provenance NFT for Chainlink VRF.', 
                a, '\\n","image":"', image, '"}'
            ));
            return concat('data:application/json;base64,', json, '', '');
        }
    }

    /// @dev Verify the token. Only callable by contract owner.
    function verify(uint[] memory _tokenIds) public onlyOwner {
        unchecked {
            for (uint i = 0; i < _tokenIds.length; i++) {
                uint tokenId = _tokenIds[i];
                tokens[tokenId].verified = true;
                emit Verified(tokenId);
            }
        }
    }
    
    /// @dev Revoke the verification. Only callable by contract owner.
    function revokeVerification(uint[] memory _tokenIds) public onlyOwner {
        unchecked {
            for (uint i = 0; i < _tokenIds.length; i++) {
                uint tokenId = _tokenIds[i];
                tokens[tokenId].verified = false;
                emit VerificationRevoked(tokenId);
            }
        }
    }

    /// @dev Returns the tokens' data in bulk.
    function tokenData(uint[] memory _tokenIds) public view returns (string[] memory) {
        unchecked {
            uint tokenCount = 0;
            uint numTokens = totalSupply();
            for (uint i = 0; i < _tokenIds.length; i++) {
                uint tokenId = _tokenIds[i];
                if (tokenId < numTokens) {
                    tokenCount++;
                }
            }
            string[] memory result = new string[](tokenCount * 12);
            uint o = 0;
            for (uint i = 0; i < tokenCount; i++) {
                uint tokenId = _tokenIds[i];
                if (tokenId < numTokens) {
                    Token memory t = tokens[tokenId];
                    result[o++] = tokenId.toString();
                    result[o++] = t.name;
                    result[o++] = t.codeURI;
                    result[o++] = t.seedKey;
                    result[o++] = t.imageURI;
                    result[o++] = t.projectURI;
                    result[o++] = addressToString(t.minter);
                    result[o++] = t.seedKeyHash.toHexString();
                    result[o++] = t.codeHash.toHexString();
                    result[o++] = t.randomness.toString();
                    result[o++] = t.paid.toString();
                    result[o++] = t.verified ? "1" : "0";
                }
            }
            return result;
        }
    }

    /// @dev Returns an array of the token ids under the owner.
    function tokensOfOwner(address _owner) external view returns(uint[] memory) {
        unchecked {
            uint tokenCount = balanceOf(_owner);
            if (tokenCount == 0) {
                return new uint[](0);
            } else {
                uint[] memory result = new uint[](tokenCount);
                for (uint index = 0; index < tokenCount; index++) {
                    result[index] = tokenOfOwnerByIndex(_owner, index);
                }
                return result;
            }
        }
    }
}