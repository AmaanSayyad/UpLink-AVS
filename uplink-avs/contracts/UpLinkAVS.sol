// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract UpLinkAVS {
    struct Request {
        uint256 id;
        string endpoint;
        address requester;
        uint256 createdAt;
        bool resolved;
        string result;
    }

    struct Operator {
        address operatorAddress;
        uint256 completedRequests;
        uint256 penalties;
    }

    uint256 public requestCounter;
    mapping(uint256 => Request) public requests;
    mapping(address => Operator) public operators;
    uint256 public rewardAmount = 0.01 ether;
    uint256 public penaltyAmount = 0.005 ether;

    event RequestSubmitted(uint256 requestId, string endpoint, address requester);
    event ResultSubmitted(uint256 requestId, string result, address operator);
    event OperatorSlashed(address operator);

    modifier onlyOperator() {
        require(operators[msg.sender].operatorAddress != address(0), "Not a registered operator");
        _;
    }

    function submitRequest(string memory endpoint) external payable {
        require(msg.value >= rewardAmount, "Insufficient fee");
        requestCounter++;
        requests[requestCounter] = Request(requestCounter, endpoint, msg.sender, block.timestamp, false, "");
        emit RequestSubmitted(requestCounter, endpoint, msg.sender);
    }

    function submitResult(uint256 requestId, string memory result) external onlyOperator {
        Request storage request = requests[requestId];
        require(!request.resolved, "Request already resolved");

        request.result = result;
        request.resolved = true;

        operators[msg.sender].completedRequests++;
        payable(msg.sender).transfer(rewardAmount);

        emit ResultSubmitted(requestId, result, msg.sender);
    }

    function slashOperator(address operator) external {
        operators[operator].penalties++;
        emit OperatorSlashed(operator);
    }

    function registerOperator() external {
        require(operators[msg.sender].operatorAddress == address(0), "Operator already registered");
        operators[msg.sender] = Operator(msg.sender, 0, 0);
    }

    receive() external payable {}
}
