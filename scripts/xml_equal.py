#!/usr/bin/env python

import sys
import difflib
import xmltodict
from datadiff import diff
from lxml import etree, objectify


def normalise_dict(d):
    """
    Recursively convert dict-like object (eg OrderedDict) into plain dict.
    Sorts list values.
    """
    out = {}
    for k, v in dict(d).iteritems():
        # Go's encoding/xml doesn't write these, so we skip them
        if k in ['@xmlns', '@xmlns:xsi', '@xsi:schemaLocation']:
            continue

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


def hack_namespace(s):
    # This is a bit of a hack: Go doesn't write XML namespace declarations, so
    # we find the first <Policy>, <PolicySet>, <Request> or <Response> tag and
    # insert it if necessary.
    if '<md:' not in s:
        return s
    if 'xmlns:md=' in s:
        return s

    tagEnd = -1
    for tag in ['PolicySet', 'Policy', 'Request', 'Response']:
        search = '<' + tag + ' '
        pos = s.find(search)
        if pos != -1:
            tagEnd = pos + len(search)
            break

    if tagEnd == -1:
        return s

    s = s[0:tagEnd] + ' xmlns:md="http://www.medico.com/schemas/record" ' + s[tagEnd:]
    return s


def main():
    with open(sys.argv[1], 'rb') as f:
        xml1 = f.read()
    with open(sys.argv[2], 'rb') as f:
        xml2 = f.read()

    # Hackishly solve the xmlns problem
    xml1 = hack_namespace(xml1)
    xml2 = hack_namespace(xml2)

    # Read and parse the input XML
    tree1 = etree.fromstring(xml1.strip())
    tree2 = etree.fromstring(xml2.strip())

    # Remove annotations and namespaces from the trees
    etree.cleanup_namespaces(tree1)
    etree.cleanup_namespaces(tree2)
    objectify.deannotate(tree1, cleanup_namespaces=True)
    objectify.deannotate(tree2, cleanup_namespaces=True)

    # Convert _back_ to a string, and then to a dict.
    cleanxml1 = etree.tostring(tree1)
    cleanxml2 = etree.tostring(tree2)

    dict1 = normalise_dict(xmltodict.parse(cleanxml1))
    dict2 = normalise_dict(xmltodict.parse(cleanxml2))

    # Finally, do the comparison
    same = dict1 == dict2

    if same:
        sys.exit(0)

    print(diff(dict1, dict2))
    sys.exit(1)


if __name__ == "__main__":
    main()
