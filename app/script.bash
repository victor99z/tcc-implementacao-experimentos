#!/bin/bash


curl -X POST http://localhost:5001/api/v1/record \
     -F 'file=@/home/jubileu/file_test' \
     -F 'json={"Pacient":"2aa850cafdf6c5d5f2d3cacd74cac8592c18024d", "Doctor": "0x2acb5cB73979eB641f0541657255578F71C87085", "Institution": "0x2Aa850cafDF6c5d5f2D3Cacd74CaC8592c18024d"}' \
     -H "Content-Type: multipart/form-data"
