Aqui está a conversão do código para Go, utilizando o algoritmo NBH (Nearest-Plane Heuristic) para adicionar uma chave unitária para cada usuário na rede de lattice:

  ```go
  package main
  
  import (
    "math/big"
  )
  
  // Lattice representa uma rede de lattice
  type Lattice struct {
    vectors [][]*big.Int
    inverse []*big.Int
    det     *big.Int
  }
  
  // NewLattice cria uma nova instância de Lattice com vetores, inverso e determinante especificados
  func NewLattice(vectors [][]*big.Int, inverse []*big.Int, det *big.Int) *Lattice {
    return &Lattice{
      vectors: vectors,
      inverse: inverse,
      det:     det,
    }
  }
  
  // AddUserKey adiciona uma chave unitária para cada usuário na rede de lattice
  func (l *Lattice) AddUserKey(userIndex int, key []*big.Int) {
    for i := range l.vectors {
      l.vectors[i][userIndex].Add(l.vectors[i][userIndex], key[i])
    }
  }
  
  // Decompose calcula a decomposição de Babai
  func (l *Lattice) Decompose(k *big.Int) []*big.Int {
    n := len(l.inverse)
    c := make([]*big.Int, n)
  
    // Calculate closest vector in lattice to <k,0,0,...> with Babai's rounding.
    for i := 0; i < n; i++ {
      c[i] = new(big.Int).Mul(k, l.inverse[i])
      round(c[i], l.det)
    }
  
    // Transform vectors according to c and subtract <k,0,0,...>.
    out := make([]*big.Int, n)
    temp := new(big.Int)
  
    for i := 0; i < n; i++ {
      out[i] = new(big.Int)
  
      for j := 0; j < n; j++ {
        temp.Mul(c[j], l.vectors[j][i])
        out[i].Add(out[i], temp)
      }
  
      out[i].Neg(out[i]).Add(out[i], new(big.Int).Mul(big.NewInt(2), l.vectors[0][i])).Add(out[i], k)
    }
  
    out[0].Add(out[0], k)
  
    return out
  }
  
  // Precompute pré-computa valores
  func (l *Lattice) Precompute(add func(i, j int)) {
    n := len(l.vectors)
    total := 1 << uint(n)
  
    for i := 0; i < n; i++ {
      for j := 0; j < total; j++ {
        if (j>>uint(i))&1 == 1 {
          add(i, j)
        }
      }
    }
  }
  
  // Multi multiplica um escalar
  func (l *Lattice) Multi(scalar *big.Int) []*big.Int {
    decomp := l.Decompose(scalar)
  
    maxLen := 0
    for _, x := range decomp {
      if len(x.Text(2)) > maxLen {
        maxLen = len(x.Text(2))
      }
    }
  
    out := make([]*big.Int, maxLen)
    for i := range out {
      out[i] = big.NewInt(0)
    }
  
    for j, x := range decomp {
      for i, bit := range x.Text(2) {
        out[i].Add(out[i], big.NewInt(0).Mul(big.NewInt(int64(bit-'0')), big.NewInt(1<<uint(j))))
      }
    }
  
    return out
  }
  
  func round(num *big.Int, denom *big.Int) {
    r := new(big.Int).Mod(num, denom)
    half := new(big.Int).Div(denom, big.NewInt(2))
    if r.Cmp(half) > 0 {
      num.Add(num, big.NewInt(1))
    }
  }
  
  func main() {
    // Suponha que tenhamos uma rede de lattice e queremos adicionar uma chave unitária para cada usuário
    // Vamos criar uma nova instância de Lattice com vetores, inverso e determinante específicos
    vectors := [][]*big.Int{
      {big.NewInt(147946756881789319000765030803803410728), big.NewInt(147946756881789319010696353538189108491)},
      {big.NewInt(147946756881789319020627676272574806254), big.NewInt(-147946756881789318990833708069417712965)},
    }
    inverse := []*big.Int{
      big.NewInt(147946756881789318990833708069417712965),
      big.NewInt(147946756881789319010696353538189108491),
    }
    det := big.NewInt(43776485743678550444492811490514550177096728800832068687396408373151616991234)
  
    lattice := NewLattice(vectors, inverse, det)
  
    // Agora, podemos adicionar uma chave unitária para cada usuário
    // Suponha que tenhamos 2 usuários e cada chave seja uma lista de valores para cada dimensão
    userKeys := [][]*big.Int{
      {big.NewInt(100000000000000000000000000000000000000), big.NewInt(100000000000000000000000000000000000000)},
      {big.NewInt(200000000000000000000000000000000000000), big.NewInt(200000000000000000000000000000000000000)},
    }
  
    // Adicionando as chaves unitárias para cada usuário na rede de lattice
    for i, key := range userKeys {
      lattice.AddUserKey(i, key)
    }
  }
  ```