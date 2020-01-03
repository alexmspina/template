import React, { useState, useEffect } from 'react'
import styled from '@emotion/styled'
import { SalesAdminServiceClient } from '../grpcweb/salesadmin_grpc_web_pb'
import {
  OrdersRequest,
  TotalSalesRevenueRequest,
  CustomerCountRequest,
  MerchantCountRequest
} from '../grpcweb/salesadmin_pb'
import SingleValue from '../widgets/singlevalue/SingleValue'
import Uploader from '../widgets/uploader/Uploader'
import Table from '../widgets/table/Table'
import Login from '../widgets/login/Login'
import { animated, useTransition } from 'react-spring'

const DashboardContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
`

const WidgetContainer = styled.div`
    display: flex;
    flex-direction; row;
    flex-wrap: wrap;
    height: 100%;
`

const Dashboard = props => {
  const { updateView, view } = props
  const [uploadSuccessful, updateUploadSuccessful] = useState(false)
  const client = new SalesAdminServiceClient('http://localhost:8080')
  console.log('client', client)

  const customerCountRequest = new CustomerCountRequest()
  const [customerCount, updateCustomerCount] = useState(0)
  useEffect(() => {
    client.getCustomerCount(customerCountRequest, {}, (err, customerCountResponse) => {
      if (err) {
        console.error('gRPC error', err)
      }
      updateCustomerCount(customerCountResponse.array[0])
      updateUploadSuccessful(false)
    })
  }, [view, uploadSuccessful])

  const merchantCountRequest = new MerchantCountRequest()
  const [merchantCount, updateMerchantCount] = useState(0)
  useEffect(() => {
    client.getMerchantCount(merchantCountRequest, {}, (err, response) => {
      if (err) {
        console.error('gRPC error', err)
      }
      updateMerchantCount(response.array[0])
      updateUploadSuccessful(false)
    })
  }, [view, uploadSuccessful])

  const ordersRequest = new OrdersRequest()
  const [orders, updateOrders] = useState([])
  useEffect(() => {
    client.getAllOrders(ordersRequest, {}, (err, response) => {
      if (err) {
        console.error('gRPC error', err)
      }
      updateOrders(response ? response.array[0] : [])
      updateUploadSuccessful(false)
    })
  }, [view, uploadSuccessful])

  const revenueRequest = new TotalSalesRevenueRequest()
  const [revenue, updateRevenue] = useState(0)
  useEffect(() => {
    client.getTotalSalesRevenue(revenueRequest, {}, (err, response) => {
      if (err) {
        console.error('gRPC error', err)
      }
      updateRevenue(response.array[0])
      updateUploadSuccessful(false)
    })
  }, [view, uploadSuccessful])

  const [widgets, setWidgets] = useState(null)
  const [loginWidget, setLoginWidget] = useState(null)

  const widgetTransitions = useTransition(widgets, null, {
    from: { transform: 'translate3d(0,1000px,0)', height: '0%' },
    enter: { transform: 'translate3d(0,0px,0)' },
    leave: { transform: 'translate3d(0,1000px,0)', height: '0%' }
  })

  const loginTransition = useTransition(loginWidget, null, {
    from: { transform: 'translate3d(0,-750px,0)', height: '0%' },
    enter: { transform: 'translate3d(0,0,0)' },
    leave: { transform: 'translate3d(0,-750px,0)', height: '0%' }
  })

  useEffect(() => {
    if (view === 'login') {
      setLoginWidget(<Login updateView={updateView} />)
      setWidgets(null)
    } else if (view === 'salesReport') {
      setLoginWidget(null)
      setWidgets([
        <WidgetContainer key='1'>
          <SingleValue title='Total Sales Revenue' value={`$${revenue ? Math.round(revenue * 100) / 100 : 0}`} />
          <SingleValue title='Customers' value={customerCount || 0} />
          <SingleValue title='Merchants' value={merchantCount || 0} />
          <Uploader grpcClient={client} updateUploadSuccessful={updateUploadSuccessful} />
          {orders.length ? <Table orders={orders} /> : null}
        </WidgetContainer>
      ])
    }
  }, [view, revenue, customerCount, merchantCount, orders])

  return (
    <DashboardContainer>
      {widgetTransitions.map(({ item, props, key }) => <animated.div key={key} style={props}>{item}</animated.div>)}
      {loginTransition.map(({ item, props, key }) => <animated.div key={key} style={props}>{item}</animated.div>)}
    </DashboardContainer>
  )
}

export default Dashboard
