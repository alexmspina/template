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
  padding: 5rem 0;
`

const WidgetContainer = styled.div`
    display: flex;
    flex-direction; row;
    flex-wrap: wrap;
    justify-content: center;
    height: 100%;
`

const Dashboard = props => {
  const { updateView, view } = props
  const [uploadSuccessful, updateUploadSuccessful] = useState(false)
  // const client = new SalesAdminServiceClient('http://localhost:8080')

  const [client, updateClient] = useState(null)
  useEffect(() => {
    const grpcClient = new SalesAdminServiceClient('/api/')
    updateClient(grpcClient)
  }, [])

  const [customerCount, updateCustomerCount] = useState(0)
  useEffect(() => {
    if (client !== null) {
      const customerCountRequest = new CustomerCountRequest()
      client.getCustomerCount(customerCountRequest, {}, (err, response) => {
        if (err) {
          console.error('customer count gRPC error', err, customerCountRequest)
        }
        console.log('customer count response', response)
        updateCustomerCount(response.array[0])
        updateUploadSuccessful(false)
      })
    }
  }, [client, view, uploadSuccessful])

  const [merchantCount, updateMerchantCount] = useState(0)
  useEffect(() => {
    if (client !== null) {
      const merchantCountRequest = new MerchantCountRequest()
      client.getMerchantCount(merchantCountRequest, {}, (err, response) => {
        if (err) {
          console.error('merchant count gRPC error', err)
        } else if (client !== null) {
          console.log('mechant count response', response)
          updateMerchantCount(response.array[0])
          updateUploadSuccessful(false)
        }
      })
    }
  }, [client, view, uploadSuccessful])

  const [orders, updateOrders] = useState([])
  useEffect(() => {
    if (client !== null) {
      const ordersRequest = new OrdersRequest()
      client.getAllOrders(ordersRequest, {}, (err, response) => {
        if (err) {
          console.error('orders gRPC error', err)
        } else if (client !== null) {
          console.log('orders response', response)
          updateOrders(response ? response.array[0] : [])
          updateUploadSuccessful(false)
        }
      })
    }
  }, [client, view, uploadSuccessful])

  const [revenue, updateRevenue] = useState(0)
  useEffect(() => {
    if (client !== null) {
      const revenueRequest = new TotalSalesRevenueRequest()
      client.getTotalSalesRevenue(revenueRequest, {}, (err, response) => {
        if (err) {
          console.error('revenue gRPC error', err)
        } else if (client !== null) {
          console.log('revenue response', response)
          updateRevenue(response.array[0])
          updateUploadSuccessful(false)
        }
      })
    }
  }, [client, view, uploadSuccessful])

  const [widgets, setWidgets] = useState(null)
  const [loginWidget, setLoginWidget] = useState(null)

  const widgetTransitions = useTransition(widgets, null, {
    config: { friction: 35 },
    from: { transform: 'translate3d(0,1000px,0)', height: '0%' },
    enter: { transform: 'translate3d(0,0px,0)' },
    leave: { transform: 'translate3d(0,1000px,0)', height: '0%' }
  })

  const loginTransition = useTransition(loginWidget, null, {
    config: { friction: 35 },
    from: { transform: 'translate3d(0,-750px,0)', height: '0%' },
    enter: { transform: 'translate3d(0,0,0)' },
    leave: { transform: 'translate3d(0,-750px,0)', height: '0%' }
  })

  useEffect(() => {
    if (view === 'login') {
      setLoginWidget(<WidgetContainer><Login updateView={updateView} /></WidgetContainer>)
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
  }, [client, updateView, view, revenue, customerCount, merchantCount, orders])

  return (
    <DashboardContainer onWheelCapture={e => e.deltaY > 0 && view !== 'salesReport' ? updateView('salesReport') : e.deltaY < 0 && view !== 'login' ? updateView('login') : null}>
      {widgetTransitions.map(({ item, props, key }) => <animated.div key={key} style={props}>{item}</animated.div>)}
      {loginTransition.map(({ item, props, key }) => <animated.div key={key} style={props}>{item}</animated.div>)}
    </DashboardContainer>
  )
}

export default Dashboard
