import React from 'react'
import styled from '@emotion/styled'
import Widget from '../Widget'

const Title = styled.div`
  font-size: 1.5rem;
  margin: 1rem 0;
`

const Value = styled.div`
  font-size: 2.5rem;
  margin: 1rem 0;
`

const SingleValue = props => {
  const { title, value } = props
  return (
    <Widget>
      <Value>{value}</Value>
      <Title>{title}</Title>
    </Widget>
  )
}

export default SingleValue
