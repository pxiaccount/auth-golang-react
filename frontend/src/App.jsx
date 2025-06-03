import { useState } from "react"
import "./App.css"

function App() {
  const [registerError, setRegisterError] = useState("")
  const [registerSuccess, setRegisterSuccess] = useState("")
  const [loginError, setLoginError] = useState("")
  const [loginSuccess, setLoginSuccess] = useState("")

  const handleRegister = async (e) => {
    e.preventDefault()
    setRegisterError("")
    setRegisterSuccess("")

    const formData = {
      Name: document.getElementById('registerName').value,
      Password: document.getElementById('registerPassword').value
    }

    try {
      const response = await fetch('http://localhost:3000/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
      })
      const data = await response.json()

      if (!response.ok) {
        setRegisterError(data.message)
        return
      }

      setRegisterSuccess(data.message)
    } catch (error) {
      setRegisterError("Failed to connect to server")
    }
  }

  const handleLogin = async (e) => {
    e.preventDefault()
    setLoginError("")
    setLoginSuccess("")

    const formData = {
      Name: document.getElementById('loginName').value,
      Password: document.getElementById('loginPassword').value
    }

    try {
      const response = await fetch('http://localhost:3000/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
      })
      const data = await response.json()

      if (!response.ok) {
        setLoginError(data.message)
        return
      }

      setLoginSuccess(data.message)
    } catch (error) {
      setLoginError("Failed to connect to server")
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center gap-8 bg-gray-50 p-4">
      <div className="w-full max-w-md">
        <form onSubmit={handleRegister} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <h1 className="text-2xl font-bold mb-4">Register</h1>
          {registerError && <div className="text-red-500 mb-4">{registerError}</div>}
          {registerSuccess && <div className="text-green-500 mb-4">{registerSuccess}</div>}

          <div className="mb-4">
            <input
              type="text"
              placeholder="Username"
              id="registerName"
              className="border-2 border-black w-full p-2 rounded"
              required
            />
          </div>

          <div className="mb-4">
            <input
              type="password"
              placeholder="Password"
              id="registerPassword"
              className="border-2 border-black w-full p-2 rounded"
              required
            />
          </div>

          <button
            type="submit"
            className="bg-blue-500 rounded-xl p-2 text-white w-full hover:bg-blue-600"
          >
            Register
          </button>
        </form>
      </div>

      <div className="w-full max-w-md">
        <form onSubmit={handleLogin} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <h1 className="text-2xl font-bold mb-4">Login</h1>
          {loginError && <div className="text-red-500 mb-4">{loginError}</div>}
          {loginSuccess && <div className="text-green-500 mb-4">{loginSuccess}</div>}

          <div className="mb-4">
            <input
              type="text"
              placeholder="Username"
              id="loginName"
              className="border-2 border-black w-full p-2 rounded"
              required
            />
          </div>

          <div className="mb-4">
            <input
              type="password"
              placeholder="Password"
              id="loginPassword"
              className="border-2 border-black w-full p-2 rounded"
              required
            />
          </div>

          <button
            type="submit"
            className="bg-blue-500 rounded-xl p-2 text-white w-full hover:bg-blue-600"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  )
}

export default App