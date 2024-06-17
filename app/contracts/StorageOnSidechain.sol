// SPDX-License-Identifier: GPL-3.0
pragma solidity >0.7.0 < 0.9.0;

contract StorageOnSidechain {

    struct Record {
        address doctor; // Token médico
        address recipient; // Token paciente
        address institution; // Token instituição
        string cid; // Dados do paciente (arquivo ou outra fonte de dados)
        uint create_at; // hora da inserção da transação
    }

    // Patient's records
    mapping(address => Record[]) records;
    mapping(address => uint) recordCounter;

    event Save(address indexed _doctor, address indexed _institution, address indexed _recipient,  uint create_at);

    function save(address _doctor, address _recipient, address _institution, string memory _data) public {
        uint create_at = block.timestamp;
        Record memory record = Record(_doctor, _recipient, _institution, _data, create_at);
        records[_recipient].push(record);

        recordCounter[_recipient] +=  1;

        emit Save(_doctor, _institution, _recipient, create_at);
    }

    function getAllRecords(address _recipient) public view returns (Record[] memory) {
        return records[_recipient];
    }

    function getRecordCounter(address _recipient) public view returns (uint) {
        return recordCounter[_recipient];
    }

    function getRecord(address _recipient, uint _index) public view returns (Record memory) {
        return records[_recipient][_index];
    }

}