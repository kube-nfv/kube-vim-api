#!/usr/bin/env python3
import json
import sys

def fix_quantity_refs(obj):
    if isinstance(obj, dict):
        if '$ref' in obj and obj['$ref'] == '#/definitions/resourceQuantity':
            return {
                'type': 'string',
                'description': obj.get('description', ''),
                'example': obj.get('example', ''),
                'pattern': '^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$'
            }
        return {k: fix_quantity_refs(v) for k, v in obj.items()}
    elif isinstance(obj, list):
        return [fix_quantity_refs(item) for item in obj]
    return obj

if len(sys.argv) != 2:
    print(f"Usage: {sys.argv[0]} <openapi-file>")
    sys.exit(1)

openapi_file = sys.argv[1]

with open(openapi_file, 'r') as f:
    spec = json.load(f)

spec = fix_quantity_refs(spec)

if 'definitions' in spec and 'resourceQuantity' in spec['definitions']:
    del spec['definitions']['resourceQuantity']

with open(openapi_file, 'w') as f:
    json.dump(spec, f, indent=2)

print(f"Fixed Quantity fields in {openapi_file}")
