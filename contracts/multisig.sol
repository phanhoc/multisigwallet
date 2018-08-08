pragma solidity ^0.4.24;

contract MultisigWallet{

    uint public nonce;                 // (only) mutable state
    uint public threshold;             // immutable state
    mapping(address => bool) isOwner; // immutable state
    address[] public owners;        // immutable state

    constructor(address[] owners_, uint threshold_) public {
        require(owners_.length >2);
        require(owners_.length <= 10 && threshold_ <= owners_.length && threshold_ > 0);

        for (uint i = 0; i < owners_.length; i++) {
            isOwner[owners_[i]] = true;
        }
        owners = owners_;
        threshold = threshold_;
    }

    // Note that address recovered from signatures must be strictly increasing, in order to prevent duplicates
    function execute(bytes signature1,bytes signature2, address destination, uint value, bytes data) public {
        var operationHash = keccak256(destination, value, data, nonce);
        // for (uint i = 0; i < signatures.length; i++) {
        //     var otherSigner = verifyMultiSig(destination, operationHash, signatures[i]);
        // }
        var otherSigner = verifyMultiSig(destination, operationHash, signature1);
        otherSigner = verifyMultiSig(destination, operationHash, signature2);

        // Success, send the transaction
        if (!(destination.call.value(value)(data))) {
            // Failed executing transaction
            revert();
        }
        nonce = nonce + 1;
    }

    function verifyMultiSig(
        address toAddress,
        bytes32 operationHash,
        bytes signature
    ) private returns (address) {

        var otherSigner = recoverAddressFromSignature(operationHash, signature);

        // Verify if we are in safe mode. In safe mode, the wallet can only send to signers
        if (!isOwner[otherSigner]) {
            // We are in safe mode and the toAddress is not a signer. Disallow!
            revert();
        }

        return otherSigner;
    }

    function recoverAddressFromSignature(
        bytes32 operationHash,
        bytes signature
    ) private pure returns (address) {
        if (signature.length != 65) {
            revert();
        }
        // We need to unpack the signature, which is given as an array of 65 bytes (like eth.sign)
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := and(mload(add(signature, 65)), 255)
        }
        if (v < 27) {
            v += 27; // Ethereum versions are 27 or 28 as opposed to 0 or 1 which is submitted by some signing libs
        }
        return ecrecover(operationHash, v, r, s);
    }

    function() payable public {}
}