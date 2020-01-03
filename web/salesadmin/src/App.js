import React, { useState } from 'react'
import styled from '@emotion/styled'
import Dashboard from './dashboard/Dashboard'
import Menu from './widgets/menu/Menu'

const StyledApp = styled.div`
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
`

const Main = styled.div`
  display: flex;
  flex-direction: row;
`

const Sidebar = styled.div`
  display: flex;
  flex-flow: column nowrap;
  min-width: 12.5%;
`

const App = () => {
  // <Sidebar updateView={updateView} />
  const [view, updateView] = useState('login')
  return (
    <StyledApp>
      <Main>
        <Sidebar>
          <Menu view={view} updateView={updateView} />
        </Sidebar>
        <Dashboard view={view} updateView={updateView} />
      </Main>
    </StyledApp>
  )
}

export default App
