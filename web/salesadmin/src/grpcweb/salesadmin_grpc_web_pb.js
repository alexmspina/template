/**
 * @fileoverview gRPC-Web generated client stub for salesadmin
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */


const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.salesadmin = require('./salesadmin_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.salesadmin.SalesAdminServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.salesadmin.SalesAdminServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.salesadmin.FileUploadRequest,
 *   !proto.salesadmin.FileUploadResponse>}
 */
const methodDescriptor_SalesAdminService_FileUpload = new grpc.web.MethodDescriptor(
  '/salesadmin.SalesAdminService/FileUpload',
  grpc.web.MethodType.UNARY,
  proto.salesadmin.FileUploadRequest,
  proto.salesadmin.FileUploadResponse,
  /**
   * @param {!proto.salesadmin.FileUploadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.FileUploadResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.salesadmin.FileUploadRequest,
 *   !proto.salesadmin.FileUploadResponse>}
 */
const methodInfo_SalesAdminService_FileUpload = new grpc.web.AbstractClientBase.MethodInfo(
  proto.salesadmin.FileUploadResponse,
  /**
   * @param {!proto.salesadmin.FileUploadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.FileUploadResponse.deserializeBinary
);


/**
 * @param {!proto.salesadmin.FileUploadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.salesadmin.FileUploadResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.salesadmin.FileUploadResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.salesadmin.SalesAdminServiceClient.prototype.fileUpload =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/salesadmin.SalesAdminService/FileUpload',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_FileUpload,
      callback);
};


/**
 * @param {!proto.salesadmin.FileUploadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.salesadmin.FileUploadResponse>}
 *     A native promise that resolves to the response
 */
proto.salesadmin.SalesAdminServicePromiseClient.prototype.fileUpload =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/salesadmin.SalesAdminService/FileUpload',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_FileUpload);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.salesadmin.OrdersRequest,
 *   !proto.salesadmin.OrdersResponse>}
 */
const methodDescriptor_SalesAdminService_GetAllOrders = new grpc.web.MethodDescriptor(
  '/salesadmin.SalesAdminService/GetAllOrders',
  grpc.web.MethodType.UNARY,
  proto.salesadmin.OrdersRequest,
  proto.salesadmin.OrdersResponse,
  /**
   * @param {!proto.salesadmin.OrdersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.OrdersResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.salesadmin.OrdersRequest,
 *   !proto.salesadmin.OrdersResponse>}
 */
const methodInfo_SalesAdminService_GetAllOrders = new grpc.web.AbstractClientBase.MethodInfo(
  proto.salesadmin.OrdersResponse,
  /**
   * @param {!proto.salesadmin.OrdersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.OrdersResponse.deserializeBinary
);


/**
 * @param {!proto.salesadmin.OrdersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.salesadmin.OrdersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.salesadmin.OrdersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.salesadmin.SalesAdminServiceClient.prototype.getAllOrders =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/salesadmin.SalesAdminService/GetAllOrders',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_GetAllOrders,
      callback);
};


/**
 * @param {!proto.salesadmin.OrdersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.salesadmin.OrdersResponse>}
 *     A native promise that resolves to the response
 */
proto.salesadmin.SalesAdminServicePromiseClient.prototype.getAllOrders =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/salesadmin.SalesAdminService/GetAllOrders',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_GetAllOrders);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.salesadmin.TotalSalesRevenueRequest,
 *   !proto.salesadmin.TotalSalesRevenueResponse>}
 */
const methodDescriptor_SalesAdminService_GetTotalSalesRevenue = new grpc.web.MethodDescriptor(
  '/salesadmin.SalesAdminService/GetTotalSalesRevenue',
  grpc.web.MethodType.UNARY,
  proto.salesadmin.TotalSalesRevenueRequest,
  proto.salesadmin.TotalSalesRevenueResponse,
  /**
   * @param {!proto.salesadmin.TotalSalesRevenueRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.TotalSalesRevenueResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.salesadmin.TotalSalesRevenueRequest,
 *   !proto.salesadmin.TotalSalesRevenueResponse>}
 */
const methodInfo_SalesAdminService_GetTotalSalesRevenue = new grpc.web.AbstractClientBase.MethodInfo(
  proto.salesadmin.TotalSalesRevenueResponse,
  /**
   * @param {!proto.salesadmin.TotalSalesRevenueRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.TotalSalesRevenueResponse.deserializeBinary
);


/**
 * @param {!proto.salesadmin.TotalSalesRevenueRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.salesadmin.TotalSalesRevenueResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.salesadmin.TotalSalesRevenueResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.salesadmin.SalesAdminServiceClient.prototype.getTotalSalesRevenue =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/salesadmin.SalesAdminService/GetTotalSalesRevenue',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_GetTotalSalesRevenue,
      callback);
};


/**
 * @param {!proto.salesadmin.TotalSalesRevenueRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.salesadmin.TotalSalesRevenueResponse>}
 *     A native promise that resolves to the response
 */
proto.salesadmin.SalesAdminServicePromiseClient.prototype.getTotalSalesRevenue =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/salesadmin.SalesAdminService/GetTotalSalesRevenue',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_GetTotalSalesRevenue);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.salesadmin.CustomerCountRequest,
 *   !proto.salesadmin.CustomerCountResponse>}
 */
const methodDescriptor_SalesAdminService_GetCustomerCount = new grpc.web.MethodDescriptor(
  '/salesadmin.SalesAdminService/GetCustomerCount',
  grpc.web.MethodType.UNARY,
  proto.salesadmin.CustomerCountRequest,
  proto.salesadmin.CustomerCountResponse,
  /**
   * @param {!proto.salesadmin.CustomerCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.CustomerCountResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.salesadmin.CustomerCountRequest,
 *   !proto.salesadmin.CustomerCountResponse>}
 */
const methodInfo_SalesAdminService_GetCustomerCount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.salesadmin.CustomerCountResponse,
  /**
   * @param {!proto.salesadmin.CustomerCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.CustomerCountResponse.deserializeBinary
);


/**
 * @param {!proto.salesadmin.CustomerCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.salesadmin.CustomerCountResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.salesadmin.CustomerCountResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.salesadmin.SalesAdminServiceClient.prototype.getCustomerCount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/salesadmin.SalesAdminService/GetCustomerCount',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_GetCustomerCount,
      callback);
};


/**
 * @param {!proto.salesadmin.CustomerCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.salesadmin.CustomerCountResponse>}
 *     A native promise that resolves to the response
 */
proto.salesadmin.SalesAdminServicePromiseClient.prototype.getCustomerCount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/salesadmin.SalesAdminService/GetCustomerCount',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_GetCustomerCount);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.salesadmin.MerchantCountRequest,
 *   !proto.salesadmin.MerchantCountResponse>}
 */
const methodDescriptor_SalesAdminService_GetMerchantCount = new grpc.web.MethodDescriptor(
  '/salesadmin.SalesAdminService/GetMerchantCount',
  grpc.web.MethodType.UNARY,
  proto.salesadmin.MerchantCountRequest,
  proto.salesadmin.MerchantCountResponse,
  /**
   * @param {!proto.salesadmin.MerchantCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.MerchantCountResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.salesadmin.MerchantCountRequest,
 *   !proto.salesadmin.MerchantCountResponse>}
 */
const methodInfo_SalesAdminService_GetMerchantCount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.salesadmin.MerchantCountResponse,
  /**
   * @param {!proto.salesadmin.MerchantCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.salesadmin.MerchantCountResponse.deserializeBinary
);


/**
 * @param {!proto.salesadmin.MerchantCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.salesadmin.MerchantCountResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.salesadmin.MerchantCountResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.salesadmin.SalesAdminServiceClient.prototype.getMerchantCount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/salesadmin.SalesAdminService/GetMerchantCount',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_GetMerchantCount,
      callback);
};


/**
 * @param {!proto.salesadmin.MerchantCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.salesadmin.MerchantCountResponse>}
 *     A native promise that resolves to the response
 */
proto.salesadmin.SalesAdminServicePromiseClient.prototype.getMerchantCount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/salesadmin.SalesAdminService/GetMerchantCount',
      request,
      metadata || {},
      methodDescriptor_SalesAdminService_GetMerchantCount);
};


module.exports = proto.salesadmin;

