# Block-Cipher-WebApp

To run, either use:

docker-compose up -d

OR

docker build -t block-cipher-webapp .
and then run it on Docker Desktop

NEW:
Just py main.py in python_src folder. Then insert ciphertext/plaintext, mode of operation, and key.

### TODO

- &#9745; Confusion (untested) & Diffusion (untested)
  - &#9744; Waktu enkripsi dan dekripsi untuk pesan dengan berbagai macam pesan (teks, file)
  - &#9744; Analisis efek longsoran (avalanche effect), yaitu bagaimana perubahan
  cipherteks jika satu bit atau satu byte plainteks atau kunci diubah
  - &#9744; Analisis ruang kunci (key space)
  - &#9744; Analisis keamanan lainnya
- &#9745; 128 bit / 16 byte blocks
- &#9745; 128 bit - 256 bit key
- &#9745; Feistel
  - &#9745; Round function
    > $plaintext: P$ <br>
    > $ciphertext: C$ <br>
    > $blocksize: n = 128$ <br>
    > $blocks: l = \lceil \dfrac{|P|}{n} \rceil$
    >
    > $block:$ $$B_i = (P_{(i - 1) \times n} \cdots P_{i \times n})$$
    >
    > $\newcommand\doubleplus{+\kern-1.5ex+\kern+0ex}$
    > $function: $ $$f(B_i) = p(s(B_i))$$
    > $$\overset{l}{\underset{i=0}\doubleplus} f(B_i) \equiv C$$
    > 
    > $legend:$ <br>
    > $\bullet {\doubleplus} \text{, is the aggregated concatenation operator similar to}\sum\text{, inspired by Haskell}$
    - &#9745; Substitution <br>
      > $blocklen: j = |B_i|$ <br>
      > $byte: b_n \in B_{i}$ <br>
      > $block: B_i = \lbrace b_0 \cdots b_j \rbrace$ <br>
      > 
      > $subs:$ $$s(B_i) = \overset{j}{\underset{n=0}\doubleplus} sBox(b_n),$$
      > where sBox is a precalculated substitution matrix
    - &#9745; Permutation <br>
      > $perm:$ $$p(B_i) = B_i \lll 2 \times ( i + 1 ) $$
      > $legend$ <br>
      > $\bullet \lll$, is a rotational left bitwise shift
  - &#9745; Repeated Cipher, 10 - 16 rounds
- &#9744; Substitution (Not sure ini ada bedanya ngga sama yg diatas)
- &#9744; Transposition 
- &#9744; Modes: 
  - &#9744; ECB
  - &#9744; CBC
  - &#9744; OFB
  - &#9744; CFB
  - &#9744; CTR
- &#9744; [Live Demo](goblc.nathancs.dev)