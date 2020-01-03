import React, { useState } from 'react'
import styled from '@emotion/styled'

const Widget = styled.div`
  display: flex;
  flex-flow: column;
  align-items: flex-start;
  color: #4f4f4f;
  height: 30rem;
  width: 100%;
  font-size: 2rem;
  margin: 10rem 0 0 -3rem;
`

const Title = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  margin: 1rem 0.5rem;
`

const Input = styled.input`
    border-top-style: hidden;
    border-right-style: hidden;
    border-left-style: hidden;
    border-bottom-style: hidden;
    background-color: #bfbfbf;
    border-radius: 3px;
    padding: 0.5rem;
    margin: 0.5rem;
    font-size: 1rem;
    color: #4f4f4f;
    ::placeholder {
        color: #4f4f4f;
    }
`

const Button = styled.div`
  height: fit-content;
  white-space: nowrap;
  border-radius: 3px;
  background-color: #bfbfbf;
  cursor: pointer;
  padding: 0.5rem;
  font-size: 1rem;
  margin: 0.5rem 0.5rem;
`

const Login = props => {
  const { updateView } = props
  const [show, setShow] = useState(false)
  return (
    <Widget>
      <Title>Login</Title>
      <Input type='username' placeholder='username' />
      <Input type={show ? 'text' : 'password'} placeholder='password' />
      <Button onClick={() => setShow(!show)}>{show ? 'HidePassword' : 'Show Password'}</Button>
      <Button onClick={() => updateView('salesReport')}>Submit</Button>
    </Widget>
  )
}

export default Login
