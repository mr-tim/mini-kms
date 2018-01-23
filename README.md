Mini-Kms
========

A toy implementation of the [Hadoop KMS API](https://hadoop.apache.org/docs/stable/hadoop-kms/index.html#KMS_HTTP_REST_API), written in Go.

APIs currently implemented:
- [ ] Key creation (need to implement material generation)
- [ ] Rollover Key
- [ ] Delete Key
- [ ] Get Key Metadata
- [ ] Get Current Key 
- [ ] Generate Encrypted Key
- [ ] Decrypt Encrypted Key
- [ ] Get Key Version
- [ ] Get Key Versions
- [ ] Get Key Names
- [ ] Get Keys Metadata

Things that still need doing
- [ ] Tests
- [ ] ACLs/security
- [ ] Check performance/concurrency
- [ ] Tidy up error handling - reporting user vs server errors