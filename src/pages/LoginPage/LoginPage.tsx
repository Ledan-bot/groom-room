import React from 'react'
import Header from '../../components/header/Header'
import Login from '../../components/login/Login'
import NewAccount from '../../components/new_account/NewAccount'
import './LoginPage.css'

const  LoginPage: React.FC = () => {
  return (
    <>
      <Header/>
      <main>
        <Login />
        <NewAccount />
      </main>
    </>
  )
}

export default LoginPage