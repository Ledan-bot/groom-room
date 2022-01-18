import  Axios  from 'axios'
import React, {SyntheticEvent, useState} from 'react'

const Login: React.FC = () => {
  let [email, updateEmail] = useState("")
  let [password, updatePassword] = useState("")

  const handleSubmit = (e: SyntheticEvent) => {
    e.preventDefault()
    if (!checkEmail(email)) {
      alert("Incorrect Email Format")
    }
    const body: {
      email: string,
      password: string
    } = {
      email: email,
      password: password
    }
    Axios.post("/api/users/login", body)
      .then(res => {
        console.log(res.data.status)
      })
  }

  const checkEmail = (email: string): boolean => {
    const regex = /[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/i
    return (regex.test(email))
  }
  return (
    <section>
      <form onSubmit={handleSubmit}>
        <h1>Login</h1>
        <input type="text" name="login-email" placeholder="Enter Email" value={email} onChange={(e) => updateEmail(e.target.value)}/>
        <input type="password" name="login-password" placeholder="Enter Password" value={password} onChange={e => updatePassword(e.target.value)}/>
        <button type="submit">Login</button>
      </form>
    </section>
  )
}

export default Login