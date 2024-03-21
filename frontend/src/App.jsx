import gopherLogo from '/logo.png'
import './App.css'

function App() {
  // Encrypt
  fetch('http://localhost:8080/encrypt', {
    method: 'POST',
    body: JSON.stringify({ 
      message: 'INSERT PLAINTEXT',
      key: 'INSERT KEY',
      mode: 'INSERT MODE',
    }),
  })
    .then((response) => response.json())
    .then((data) => console.log("Results from encrypt:", data))

  // Decrypt
  fetch('http://localhost:8080/decrypt', {
    method: 'POST',
    body: JSON.stringify({ 
      message: 'INSERT CIPHERTEXT',
      key: 'INSERT KEY',
      mode: 'INSERT MODE',
    }),
  })
    .then((response) => response.json())
    .then((data) => console.log("Results from decrypt:", data))

  return (
    <>
      <div>
        <a href="#" target="_blank">
          <img src={gopherLogo} className="logo" alt="Gopher logo" />
        </a>
      </div>
      <h1>GoBlockC</h1>
      <p className="read-the-docs">
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Dolorum consectetur corporis quos delectus, quam magni, doloremque in sunt eum cumque inventore sapiente obcaecati, hic nostrum similique. Porro aliquid sunt voluptatem?
      </p>
    </>
  )
}

export default App
