import React, {SyntheticEvent, useState} from 'react'

const Login: React.FC = () => {
  let [email, updateEmail] = useState("")
  let [password, updatePassword] = useState("")

  const handleSubmit = (e: SyntheticEvent) => {
    e.preventDefault()
    alert("hello")
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