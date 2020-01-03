import React from 'react'
import styled from '@emotion/styled'

const StyledTable = styled.table`
    color: #4f4f4f;
    border: none;
    margin: 1.6rem;
`

const TableHeader = styled.thead`
    background-color: #b3b3b396;
`

const OddRow = styled.tr`
    background-color: #d2d0d296;
`

const EvenRow = styled.tr`
    background-color: #c4c0c496;
`

const Cell = styled.div`
    display: inline-flex;
    justify-content: center;
    padding: 1rem;
`

const HeaderCell = styled.div`
    display: inline-flex;
    justify-content: center;
    padding: 1rem;
`

const Table = props => {
  const { orders } = props
  const rows = orders.map((order, i) => i % 2 === 0 ? <OddRow key={Math.random()}>{order.map(cell => <td key={Math.random()}><Cell>{cell}</Cell></td>)}</OddRow> : <EvenRow key={Math.random()}>{order.map(cell => <td key={Math.random()}><Cell>{cell}</Cell></td>)}</EvenRow>)
  return (
    <StyledTable cellSpacing='0' cellPadding='0'>
      <TableHeader>
        <tr>
          <th colSpan='1'><HeaderCell>Order ID</HeaderCell></th>
          <th colSpan='1'><HeaderCell>Customer</HeaderCell></th>
          <th colSpan='1'><HeaderCell>Item</HeaderCell></th>
          <th colSpan='1'><HeaderCell>Price</HeaderCell></th>
          <th colSpan='1'><HeaderCell>Quantity</HeaderCell></th>
          <th colSpan='1'><HeaderCell>Merchant</HeaderCell></th>
          <th colSpan='1'><HeaderCell>Merchant Address</HeaderCell></th>
        </tr>
      </TableHeader>
      <tbody>{rows}</tbody>
    </StyledTable>
  )
}

export default Table
