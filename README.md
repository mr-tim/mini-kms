Mini-Kms
========

A toy implementation of the [Hadoop KMS API](https://hadoop.apache.org/docs/stable/hadoop-kms/index.html#KMS_HTTP_REST_API), written in Go.

APIs currently implemented:
- [x] Key creation
- [x] Rollover Key
- [x] Delete Key
- [x] Get Key Metadata
- [x] Get Current Key 
- [ ] Generate Encrypted Key
- [ ] Decrypt Encrypted Key
- [x] Get Key Version
- [x] Get Key Versions
- [x] Get Key Names
- [x] Get Keys Metadata

Things that still need doing
- [ ] Tests
- [ ] ACLs/security
- [ ] Check performance/concurrency
- [ ] Tidy up error handling - reporting user vs server errors