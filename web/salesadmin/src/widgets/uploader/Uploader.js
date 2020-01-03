import React, { useCallback } from 'react'
import { useDropzone } from 'react-dropzone'
import styled from '@emotion/styled'
import {
  File,
  FileUploadRequest
} from '../../grpcweb/salesadmin_pb'
import Widget from '../Widget'

const UploadMessage = styled.div`
  font-size: 1.5rem;
  margin: 1rem 0;
  pointer-events: none;
`

const Uploader = props => {
  const { grpcClient, updateUploadSuccessful } = props

  const onDrop = useCallback(acceptedFiles => {
    acceptedFiles.forEach((file) => {
      const reader = new FileReader()
      const salesFile = new File()
      reader.readAsBinaryString(file)
      reader.onabort = () => console.log('file reading was aborted')
      reader.onerror = () => console.log('file reading has failed')
      reader.onload = () => {
        const fileString = reader.result
        const fileStringRows = fileString.split('\n')
        const uint8matrix = fileStringRows.map(row => {
          const uint8Row = new Uint8Array([...Buffer.from(row)])
          return uint8Row
        })
        salesFile.setFileName(file.path)
        salesFile.setFileBytesList(uint8matrix)
        const fileUploadRequest = new FileUploadRequest()
        fileUploadRequest.setFile(salesFile)
        grpcClient.fileUpload(fileUploadRequest, {}, (err, response) => {
          if (err) {
            console.error('gRPC error', err)
          }
          updateUploadSuccessful(response.array[0])
        })
      }
    })
  }, [])

  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop })
  return (
    <Widget className='widget-singlevalue'>
      <div {...getRootProps()}>
        <input {...getInputProps()} />
        {
          isDragActive
            ? <UploadMessage>Release sale file to upload</UploadMessage>
            : <UploadMessage>Drop a salesdata.csv here</UploadMessage>
        }
      </div>
    </Widget>
  )
}

export default Uploader
