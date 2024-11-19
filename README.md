# **UpLink-AVS**

## **Overview**
UpLink-AVS is a minimal Active Validator Service (AVS) designed to determine probabilistically whether an endpoint (server, URL, or IP) is up/reachable and confirm its status on-chain. Built with reusability and extensibility in mind, this project serves as an educational and foundational framework for developers working within the EigenLayer ecosystem.

## **Motivation**
1. **Reusability**: Provide a modular AVS that can be extended or integrated into other AVSs.
2. **Education**: Demonstrate the minimum capabilities required for AVS participation in EigenLayer.
3. **Transparency**: Record endpoint liveliness on-chain for verifiable and immutable status logs.

## **Use Case**
An AVS service operates an "active set" of node operators to monitor endpoint responsiveness. UpLink-AVS maintains an on-chain record of which operators are truly responsive and penalizes those who fail to report or submit incorrect results.

---

## **Process Flow Architecture**
![Screenshot 2024-11-19 090647](https://github.com/user-attachments/assets/2092ecb7-8df0-4cac-b2f8-d5005b6d4ac3)


1. **Request Submission**:
   - Users submit a request to test the liveliness of an endpoint.
   - The request is logged on-chain via the UpLink-AVS smart contract.
2. **Operator Validation**:
   - A network of operators retrieves the request and tests the endpoint's status using ping and traceroute.
   - Results are submitted back to the smart contract.
3. **Incentives**:
   - Responsive operators are rewarded.
   - Non-responsive or malicious operators are slashed.
4. **Recordkeeping**:
   - Results are aggregated and recorded on-chain for transparency.

---

## **Features**
- **Ping Test**: Verifies endpoint reachability.
- **Traceroute Test**: Analyzes the network path to the endpoint.
- **On-Chain Record**: Logs endpoint status for verifiable tracking.
- **Incentive Mechanism**:
  - Rewards for valid results.
  - Slashing for non-compliance or malicious behavior.

---

## **Special Thanks**
I would like to thank the following individuals and organizations for their support:

- **Wes Floyd, Aarav Raj, Shan Rasool:** For providing guidance and inspiration that helped shape the vision of UpLink-AVS.

- **Mustafa:** For technical insights and fostering a spirit of innovation.

- **CollegeDAO:** For creating a collaborative and learning-focused ecosystem.

- **Blockchain Acceleration Foundation (BAF):** For supporting initiatives that push the boundaries of decentralized technologies.

- **B@B:** For empowering developers to build impactful blockchain-based solutions.

Your contributions have been instrumental in making this project a reality. Thank you for being a part of the journey.
