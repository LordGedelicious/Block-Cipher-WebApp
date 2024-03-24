# Block-Cipher-WebApp

To run, either use:

docker-compose up -d

OR

docker build -t block-cipher-webapp .
and then run it on Docker Desktop

NEW:
Just py main.py in python_src folder. Then insert ciphertext/plaintext, mode of operation, and key.

## TODO

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
    > $bit: \beta$ <br>
    > $plaintext: P = \lbrace  \beta_0 \cdots \beta_{|P|} \rbrace$ <br>
    > $blocksize: n = 128$ <br>
    > $byte: b_x^i \in B_{i}, x \in [0 \cdots \frac{n}{8})$ <br>
    > $block:$ <br>
    > $$B_i = \lbrace  \beta_{i \times n} \cdots \beta_{(i+1) \times n} \rbrace$$
    > $$B_i = \lbrace b_0^i \cdots b_{|B_i|}^i \rbrace$$
    > $key: K\text{, }|K| = 128$ <br>
    > $subkey:$
    > $$\Bbb{K}_y = K \lll 8 \times n$$
    > $$\Bbb{K}_y = \lbrace \beta_0 \cdots \beta_{|K|} \rbrace$$
    > $$K_y = \lbrace \beta_0 \cdots \beta_{|K| \div 2} \rbrace$$
    > $lhs \_ initial:$ <br>
    > $$L_0^i = \lbrace b_0^i \cdots b_{|B_i| \div 2}^i  \rbrace$$
    > $rhs \_ initial:$ <br>
    > $$R_0^i = \lbrace  b_{|B_i| \div 2}^i \cdots b_{|B_i|}^i \rbrace$$
    > $\newcommand\doubleplus{+\kern-1.5ex+\kern+0ex}$
    > $round \_ function:$
    > $$f(B_i,j) = R_{j-1}^i \| L_{j-1}^i \otimes F(R_{j-1}^i, K_{j - 1})$$
    > $feistel \_ function:$
    > $$F(R^i_z, K_z) = P(S(R^i_z)) \otimes K_z$$
    > $blocks: l = \lceil \dfrac{|P|}{n} \rceil$ <br>
    > $rounds: r = 16$ <br>
    > $ciphertext:$
    > $$C = \overset{l}{\underset{i=0}\doubleplus} f(B_i)^{r}$$
    > $$f(B_i)^{r} = f(f(\dots{f(B_i, 1)}, r-1), r)$$
    >
    > $legend:$ <br>
    > $\bullet \lll \text{, is a rotational left bitwise shift}$ <br>
    > $\bullet \otimes \text {, is the XOR function}$ <br>
    > $\bullet {\doubleplus} \text{, is the aggregated concatenation operator similar to}\sum\text{, inspired by Haskell}$ <br>
    - &#9745; Substitution <br>
      >
      > $subs:$ $$s(B_i) = \overset{j}{\underset{n=0}\doubleplus} sBox(b_n),$$
      > where sBox is a precalculated substitution matrix
    - &#9745; Permutation <br>
      > $perm:$ $$p(B_i) = B_i \lll 2 \times ( i + 1 ) $$
      > $legend$ <br>
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
