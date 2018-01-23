Mini-Kms
========

A toy implementation of the [Hadoop KMS API](https://hadoop.apache.org/docs/stable/hadoop-kms/index.html#KMS_HTTP_REST_API), written in Go.

APIs currently implemented:
- [x] Key creation
- [ ] Rollover Key
- [ ] Delete Key
- [x] Get Key Metadata
- [ ] Get Current Key 
- [ ] Generate Encrypted Key
- [ ] Decrypt Encrypted Key
- [ ] Get Key Version
- [ ] Get Key Versions
- [x] Get Key Names
- [ ] Get Keys Metadata

Things that still need doing
- [ ] Tests
- [ ] ACLs/security
- [ ] Check performance/concurrency
- [ ] Tidy up error handling - reporting user vs server errors