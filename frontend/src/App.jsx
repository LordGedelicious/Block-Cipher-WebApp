import gopherLogo from '/logo.png'
import './App.css'

function App() {
  // request to backend, allow cors
  fetch('http://localhost:8080/encrypt', {
    method: 'GET',
  })
    .then((response) => response.json())
    .then((data) => console.log(data))

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
