#!/usr/bin/env python

import sys
import xmltodict
from lxml import etree

def normalise_dict(d):
    """
    Recursively convert dict-like object (eg OrderedDict) into plain dict.
    Sorts list values.
    """
    out = {}
    for k, v in dict(d).iteritems():
        if hasattr(v, 'iteritems'):
            out[k] = normalise_dict(v)
        elif isinstance(v, list):
            out[k] = []
            for item in sorted(v):
                if hasattr(item, 'iteritems'):
                    out[k].append(normalise_dict(item))
                else:
                    out[k].append(item)
        else:
            out[k] = v

    return out


def xml_compare(a, b):
    """
    Compares two XML documents (as string or etree)

    Does not care about element order
    """
    if not isinstance(a, basestring):
        a = etree.tostring(a)
    if not isinstance(b, basestring):
        b = etree.tostring(b)

    a = normalise_dict(xmltodict.parse(a))
    b = normalise_dict(xmltodict.parse(b))
    return a == b


def main():
    with open(sys.argv[1], 'rb') as f:
        xml1 = f.read()
    with open(sys.argv[2], 'rb') as f:
        xml2 = f.read()

    tree1 = etree.fromstring(xml1.strip())
    tree2 = etree.fromstring(xml2.strip())

    if xml_compare(tree1, tree2):
        sys.exit(0)
    else:
        sys.exit(1)


if __name__ == "__main__":
    main()
