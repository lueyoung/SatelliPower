#!/usr/bin/env python

import json
from random import randint

def get_dict1():
  src = {}

  """
  for apid 0060
  """
  src["apid_0060"] = '0' * 16		# unit_dec, 16bits: 0 - 15

  src['packetSeqId_0060'] = '00'  	# unit_dec,  2bits: 16, 17
  src['packetCount_0060'] = '0' * 14	# unit_dec, 14bits: 
  src['packetLength_0060'] = '0' * 16	# unit_dec

  src['TMN001_1'] = '0' * 16		# ployfit_unit
  src['TMN002_1'] = '0' * 16		# ployfit_unit
  src['TMN003_1'] = '0' * 16		# ployfit_unit
  src['TMN004_1'] = '0' * 16		# ployfit_unit
  src['TMN005_1'] = '0' * 16		# ployfit_unit
  src['TMN006'] = '0' * 16		# ployfit_unit
  src['TMN007'] = '0' * 16		# ployfit_unit
  src['TMN008'] = '0' * 16		# ployfit_unit

  src['TMN009'] = '0'			# unit_dec
  src['TMN010'] = '0'			# unit_dec
  src['TMN011'] = '0'			# unit_dec
  src['TMN012'] = '0'			# unit_dec
  src['TMN013'] = '0' * 4		# unit_dec

  src['TMN014'] = '0' * 8		# ployfit_unit
  src['TMN015'] = '0' * 8		# ployfit_unit
  src['TMN016'] = '0' * 8		# ployfit_unit
  src['TMN017'] = '0' * 8		# ployfit_unit
  src['TMN018'] = '0' * 8		# ployfit_unit
  src['TMN019'] = '0' * 8		# ployfit_unit

  src['TMN020'] = '0'			# unit_dec
  src['TMN021'] = '0'			# unit_dec
  src['TMN022'] = '0'			# unit_dec
  src['TMN023'] = '0'			# unit_dec
  src['TMN024'] = '0'			# unit_dec
  src['TMN025'] = '0'			# unit_dec
  src['TMN026'] = '0' * 2		# unit_dec
  src['TMN027'] = '0' * 8		# unit_dec
  src['TMN028'] = '0' * 8		# unit_dec
  src['TMN029'] = '0' * 8		# unit_dec
  src['TMN030'] = '0' * 8		# unit_dec
  src['TMN031'] = '0' * 8		# unit_dec

  return src

def get_dict2():
  src = {}

  """
  for apid 0060
  """
  src["apid_0064"] = '0' * 16		# unit_dec, 16bits: 0 - 15

  src['packetSeqId_0064'] = '00'  	# unit_dec,  2bits: 16, 17
  src['packetCount_0064'] = '0' * 14	# unit_dec, 14bits: 
  src['packetLength_0064'] = '0' * 16	# unit_dec

  src['TMN101'] = '0' * 8		# ployfit_unit
  src['TMN102'] = '0' * 8		# ployfit_unit
  src['TMN103'] = '0' * 8		# ployfit_unit
  src['TMN104'] = '0' * 8		# ployfit_unit
  src['TMN105'] = '0' * 8		# ployfit_unit
  src['TMN106'] = '0' * 8		# ployfit_unit

  src['TMN107'] = '0'			# unit_dec
  src['TMN108'] = '0'			# unit_dec
  src['TMN109'] = '0'			# unit_dec
  src['TMN110'] = '0'			# unit_dec
  src['TMN111'] = '0'			# unit_dec
  src['TMN112'] = '0' * 3		# unit_dec

  src['TMN113'] = '0' * 16		# ployfit_unit
  src['TMN114'] = '0' * 16		# ployfit_unit
  src['TMN115'] = '0' * 16		# ployfit_unit
  src['TMN116'] = '0' * 16		# ployfit_unit
  src['TMN117'] = '0' * 16		# ployfit_unit
  src['TMN118'] = '0' * 16		# ployfit_unit
  src['TMN119'] = '0' * 16		# ployfit_unit
  src['TMN120'] = '0' * 16		# ployfit_unit
  src['TMN121'] = '0' * 16		# ployfit_unit
  src['TMN122'] = '0' * 16		# ployfit_unit
  src['TMN123'] = '0' * 16		# ployfit_unit

  src['TMN124'] = '0'			# unit_dec
  src['TMN125'] = '0'			# unit_dec
  src['TMN154'] = '0'			# unit_dec
  src['TMN126'] = '0' * 5		# unit_dec

  src['TMN127'] = '0' * 8		# therm3
  src['TMN128'] = '0' * 8		# therm3
  src['TMN129'] = '0' * 8		# therm3
  src['TMN130'] = '0' * 8		# therm3
  src['TMN131'] = '0' * 8		# therm3

  src['TMN132'] = '0'			# unit_dec
  src['TMN133'] = '0' * 2		# unit_dec
  src['TMN134'] = '0' * 2		# unit_dec
  src['TMN135'] = '0' * 3		# unit_dec
  src['TMN136'] = '0' * 8		# unit_dec
  src['TMN137'] = '0' * 8		# unit_dec

  src['TMN138'] = '0' * 8		# ployfit_unit
  src['TMN139'] = '0' * 16		# ployfit_unit

  src['TMN140'] = '0' * 8		# unit_dec
  src['TMN141'] = '0' * 8		# unit_dec
  src['TMN141_b7'] = '0'		# unit_dec
  src['TMN141_b6'] = '0'		# unit_dec
  src['TMN141_b5'] = '0'		# unit_dec
  src['TMN141_b4'] = '0'		# unit_dec
  src['TMN141_b3'] = '0'		# unit_dec
  src['TMN141_b2'] = '0'		# unit_dec
  src['TMN141_b1'] = '0'		# unit_dec
  src['TMN141_b0'] = '0'		# unit_dec
  src['TMN142'] = '0' * 8		# unit_dec
  src['TMN142_b2'] = '0'		# unit_dec
  src['TMN142_b1'] = '0'		# unit_dec
  src['TMN142_b0'] = '0'		# unit_dec
  src['TMN143'] = '0' * 8		# unit_dec
  src['TMN144'] = '0' * 8		# unit_dec

  src['TMN145'] = '0' * 8		# ployfit_unit
  src['TMN146'] = '0' * 8		# ployfit_unit
  src['TMN147'] = '0' * 8		# ployfit_unit
  src['TMN148'] = '0' * 8		# ployfit_unit

  src['TMN149'] = '0' * 8		# therm3
  src['TMN150'] = '0' * 8		# therm3
  src['TMN151'] = '0' * 8		# therm3
  src['TMN152'] = '0' * 8		# therm3

  src['TMN153'] = '0' * 8		# unit_dec

  return src

def make_json(d, id=0):
    for key in d.keys():
        if key == 'apid_0060':
            d[key] = '0060'
            continue
        if key == 'apid_0064':
            d[key] = '0064'
            continue
        v = d[key]
        d[key] = str(randint(0,2 ** len(v) - 1))
    d['ID'] = str(id)
    return json.dumps(d)


def main():
  d = get_dict1()
  j = make_json(d)
  print(j)
  d = get_dict2()
  j = make_json(d)
  print(j)
  

if __name__ == "__main__":
  main()
