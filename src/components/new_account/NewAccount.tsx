import  Axios  from 'axios'
import React, { SyntheticEvent, useState } from 'react'

const NewAccount: React.FC = () => {
  let [firstName, updateFirstName] = useState('')
  let [lastName, updateLastName] = useState('')
  let [email, updateEmail] = useState('')
  let [password, updatePassword] = useState('')

  const handleSubmit = (e: SyntheticEvent) => {
    e.preventDefault
    if (!checkEmail(email)) {
      alert("Incorrect Email format")
    }

    type Body = {
      firstname: string,
      lastname: string,
      email: string,
      password: string
    }
    const body: Body = {
      firstname: firstName,
      lastname: lastName,
      email: email,
      password: password
    }
    Axios.post("/api/users/create", body)
      .then(res => {
        console.log(res)
      })
  }

  const checkEmail = (email: string): boolean => {
    const regex = /[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/i
    return (regex.test(email))
  }
  return (
    <section>
      <form onSubmit={handleSubmit}>
        <h1>Welcome!</h1>
        <p>Flease fill out all entries to creat an account!</p>
        <input type="text" placeholder="First Name" value={firstName} onChange={e => updateFirstName(e.target.value)}/>
        <input type="text" placeholder="Last Name" value={lastName} onChange={e => updateLastName(e.target.value)}/>
        <input type="text" placeholder="Enter Email" value={email} onChange={e => updateEmail(e.target.value)}/>
        <input type="password" placeholder="Enter Password" value={password} onChange={e => updatePassword(e.target.value)}/>
        <button type="submit">Create Account</button>
      </form>
    </section>
  )
}

export default NewAccount