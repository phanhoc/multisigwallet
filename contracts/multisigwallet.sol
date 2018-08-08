pragma solidity ^0.4.22;

contract MultisigWallet {

    uint public nonce;                 // (only) mutable state
    uint public threshold;             // immutable state
    mapping(address => bool) isOwner; // immutable state
    address[] public ownersArr;        // immutable state

    // Note that owners_ must be strictly increasing, in order to prevent duplicates
    constructor(address[] owners_, uint threshold_) public {
        require(owners_.length <= 10 && threshold_ <= owners_.length && threshold_ >= 0);

        for (uint i = 0; i < owners_.length; i++) {
            isOwner[owners_[i]] = true;
        }
        ownersArr = owners_;
        threshold = threshold_;
    }

    // Note that address recovered from signatures must be strictly increasing, in order to prevent duplicates
    function execute(uint8[] sigV, bytes32[] sigR, bytes32[] sigS, address destination, uint value, bytes data) public {
        require(sigR.length == threshold);
        require(sigR.length == sigS.length && sigR.length == sigV.length);

        // Follows ERC191 signature scheme: https://github.com/ethereum/EIPs/issues/191
        bytes32 txHash = keccak256(byte(0x19), byte(0), this, destination, value, data, nonce);

        // cannot have address(0) as an owner
        for (uint i = 0; i < threshold; i++) {
            address recovered = ecrecover(txHash, sigV[i], sigR[i], sigS[i]);
            require(isOwner[recovered]);
        }

        // If we make it here all signatures are accounted for.
        // The address.call() syntax is no longer recommended, see:
        // https://github.com/ethereum/solidity/issues/2884

        bool success = false;
        //        success = destination.call.value(value)(data);

        assembly {success := call(gas, destination, value, add(data, 0x20), mload(data), 0, 0)}
        require(success);
        nonce = nonce + 1;
    }

    function() payable public {}
}
