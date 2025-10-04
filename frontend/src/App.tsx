import { useState, useEffect } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { useNavigate } from 'react-router'

function App() {
  const [backendData, setBackendData] = useState<string>('')
  const [count, setCount] = useState(0)
  const navigate = useNavigate()
  // initial fetch when the app mounts
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/')
        if (!response.ok) {
          throw new Error(`HTTP Error! status: ${response.status}`)
        }
        const data = await response.json()
        // backend returns { message: '...' }
        setBackendData(data?.message ?? JSON.stringify(data))
      } catch (err) {
        console.error('Error fetching initial data:', err)
        setBackendData('Error fetching data')
      }
    }
    fetchData()
  }, [])

  // called when user clicks "Get home message"
  const handleGetHomeMessage = async () => {
    try {
      const response = await fetch('http://localhost:8080/')
      if (!response.ok) throw new Error(`HTTP Error! status: ${response.status}`)
      const data = await response.json()
      setBackendData(data?.message ?? JSON.stringify(data))
    } catch (err) {
      console.error('Error fetching home message:', err)
      setBackendData('Error fetching home message')
    }
  }

  // called when user clicks "Get other message"
  const handleGetOtherMessage = async () => {
    try {
      const response = await fetch('http://localhost:8080/other')
      if (!response.ok) throw new Error(`HTTP Error! status: ${response.status}`)
      const data = await response.json()
      setBackendData(data?.message ?? JSON.stringify(data))
    } catch (err) {
      console.error('Error fetching other message:', err)
      setBackendData('Error fetching other message')
    }
  }

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank" rel="noreferrer">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank" rel="noreferrer">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <button onClick={handleGetHomeMessage}>Get home message</button>
      <button onClick={handleGetOtherMessage}>Get other message</button>
      <button onClick={()=>{navigate('/home')}}>Home</button>
      <div className="card">
        <button onClick={() => setCount((c) => c + 1)}>
          count is {count}
        </button>
        <p>
          {backendData}
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
