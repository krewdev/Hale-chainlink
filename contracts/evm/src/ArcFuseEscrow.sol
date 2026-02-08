// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ArcFuseEscrow {
    event Deposit(address indexed seller, address indexed depositor, uint256 amount);
    event Release(address indexed seller, uint256 amount, bytes32 indexed transactionId);
    event Withdrawal(address indexed depositor, uint256 amount, string reason);
    event ContractRequirementsSet(address indexed seller, string requirements, string sellerContact);
    event DeliverySubmitted(address indexed seller, string deliveryHash, uint256 timestamp);

    function deposit(address seller) external payable;
    function setContractRequirements(address seller, string calldata requirements, string calldata sellerContact) external;
    function submitDelivery(string calldata deliveryHash) external;
    function release(address seller, bytes32 transactionId) external;
    function refund(address seller, string calldata reason) external;
    function deposits(address) external view returns (uint256);
    
    struct DepositorInfo {
        address depositor;
        uint256 amount;
    }
    
    function getDepositors(address seller) external view returns (DepositorInfo[] memory);
    function getDepositorCount(address seller) external view returns (uint256);
    function oracle() external view returns (address);
    function owner() external view returns (address);
}
