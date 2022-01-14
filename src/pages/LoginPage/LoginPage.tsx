import React from 'react'
import Login from '../../components/login/Login'
import NewAccount from '../../components/new_account/NewAccount'
import './LoginPage.css'

const  LoginPage: React.FC = () => {
  return (
    <>
      <header>
        <h1>Groom Room User Login</h1>
      </header>
      <main>
        <Login />
        <NewAccount />
      </main>
    </>
  )
}

export default LoginPage