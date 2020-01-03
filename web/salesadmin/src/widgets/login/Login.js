import React from 'react'
import styled from '@emotion/styled'
import { FiUser } from 'react-icons/fi'

const Widget = styled.div`
  display: flex;
  flex-flow: column nowrap;
  align-items: center;
  color: #4f4f4f;
  height: 100%;
  width: 100%;
  font-size: 2rem;
  margin: 10rem 0 0 -3rem;
`

const Input = styled.input`
    border-top-style: hidden;
    border-right-style: hidden;
    border-left-style: hidden;
    border-bottom-style: hidden;
    background-color: #bfbfbf;
    border-radius: 3px;
    padding: 1rem;
    margin: 0.5rem;
    color: #4f4f4f;
    ::placeholder {
        color: #4f4f4f;
    }
`

const Button = styled.div`
    border-radius: 3px;
    background-color: #bfbfbf;
    cursor: pointer;
    padding: 0.5rem;
    font-size: 1rem;
`

const Login = props => {
  const { updateView } = props
  return (
    <Widget>
      <FiUser />
      Login
      <Input type='text' placeholder='username' />
      <Input type='text' placeholder='password' />
      <Button onClick={() => updateView('salesReport')}>Submit</Button>
    </Widget>
  )
}

export default Login
