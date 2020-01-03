import React from 'react'
import styled from '@emotion/styled'
import { FiHome, FiLogIn } from 'react-icons/fi'
import { animated, useSpring } from 'react-spring'

const Widget = styled.div`
  display: flex;
  flex-flow: column nowrap;
  margin: 0 0.5rem;
  color: #4f4f4f;
`

const MenuContent = styled.div`
    display: flex;
    flex-direction: column;
    align-content: center;
    justify-content: center
    width: fit-content;
    padding: 0.75rem;
    position: fixed;
`

const MenuButtons = styled.div`
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column; 
    justify-content: center;
`

const MenuItem = styled.div`
  width: fit-content;
  cursor: pointer;
  margin: 1rem;
  font-size: 2rem;
  padding: 1rem;
  border-radius: 3px;
  color: ${props => props.selected ? '#000000' : '#4f4f4f'};
`

const MenuSelectSquare = styled.div`
  border-radius: 3px;
  background-color: #c4c0c496;
  position: absolute;
  height: 4rem;
  width: 4rem;
  margin: 1rem;
`

const Logo = styled.div`
  display: flex:
  flex-direction: column;
  justify-content: center;
  margin: 1rem;
  font-size: 2rem;
  width: 5rem;
  pointer-events: none; 
`

const Title = styled.div`
  display: inline-flex;
`

const salesReportView = 'salesReport'
const loginView = 'login'

const Menu = props => {
  const { view, updateView } = props
  const styleProps = useSpring({ transform: view === 'login' ? 'translate3d(0, 105px, 0)' : 'translate3d(0, 0px, 0)' })
  return (
    <Widget>
      <MenuContent className='sidebar'>
        <Logo><Title>Sales Admin</Title></Logo>
        <MenuButtons>
          <animated.div style={styleProps}><MenuSelectSquare /></animated.div>
          {view === salesReportView ? <MenuItem selected><FiHome /></MenuItem> : <MenuItem onClick={() => updateView(salesReportView)}><FiHome /></MenuItem>}
          {view === loginView ? <MenuItem selected><FiLogIn /></MenuItem> : <MenuItem onClick={() => updateView(loginView)}><FiLogIn /></MenuItem>}
        </MenuButtons>
      </MenuContent>

    </Widget>
  )
}

export default Menu
