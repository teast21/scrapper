# -*- coding: utf-8 -*-
"""
Ethereum deposit data scraper
@teastman
"""

import argparse
import json
from pathlib import Path

p = argparse.ArgumentParser()
p.add_argument("file_path", type=Path)
file = p.parse_args()
out = open(file.file_path)
data = json.load(out) 

parsed = json.dumps(data)
with open("test.txt", "w") as f:
    
    for i in range(len(data)):
        f.write(data[i]["pubkey"])
    f.write('\n')
    
    for i in range(len(data)):
        f.write(data[i]["withdrawal_credentials"])
    f.write('\n')
    
    for i in range(len(data)):
        f.write(data[i]["signature"])
    f.write('\n')
    
    for i in range(len(data)):
        f.write(data[i]["deposit_data_root"])
    f.write('\n')
f.close()
