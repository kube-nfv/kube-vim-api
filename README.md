API for kube-vim application.

Repository contains proto defentions of the vi-vnfm, or-vi ETSI MANO reference points descirbed in
ETSI GS NFV-IFA 006 (or-vi) [^fn1], ETSI GS NFV-IFA 005 (vi-vnfm) [^fn2]. Some models defentions was used from the
ETSI GS NFV-SOL 014.

RESTful reverse proxy generated using grpc-gateway and conforms the ETSI GS NFV-SOL 013 [^fn3].

/kube-ovn-api contains golang types to work with [kube-ovn](https://www.kube-ovn.io) as well as code autogeneration logic
for Golang ClientSet and helpers.




## References

[^fn1]:
[^fn2]:
[^fn3]: [ETSI GS NFV-SOL 013 Release v5.2.1 (2024-12)](https://www.etsi.org/deliver/etsi_gs/NFV-SOL/001_099/013/05.02.01_60/gs_nfv-sol013v050201p.pdf)
