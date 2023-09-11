  MONGODB (basics)
-------------------

* Install mongodb in your system
* Use command to start(in linux deb)
      sudo systemctl start mongodb
* Check status
      sudo systemctl status mongodb
* To access mongodb
      mongo

* Basic Ideas and concepts
   ---------------------
1.  Document-Oriented: MongoDB is a document-oriented database, which means it stores data in documents. Each document is a self-contained
    unit with its own data and schema. Documents are tipically stored in collections, which are alnaloghous to tables in relational database.
2.  BSON Format: MongoDB stores data in BSON(Binary JSON), which is a binary-encoded serialization of JSON-like documents. BSON allows for
    efficient storage and retrival of data.
3.  Collections: Collections in MongoDB are groups of documents that are similar in structure. Collections are analogous to tables in
    relational database. Documents within a collection can have different fields, but they tipically share a common purpose or type of data.
4.  Documents: Documents are individual records in MongoDB. They are represented as JSON-like objects with key-value pairs. Fields in
    can have different data types, including strings, numbers, arrays, and nested documents.
5.  ObjectId: MongoDB assigns a unique ObjectId to each document when it is created. This ObjectId serves as a primary key for the document
    within its collection.
6.  Schemaless: MongoDB is often referred to as schemaless database because documents within a collection can have different stuctures. This
    flexibility allows for easy adaptation to changing data requirements.
7.  Indexing: MongoDB supports indexing on fields within documents, which can significantly improve query performance. You can create single-
    field of compound indexes to optimize queries.
8.  Query Language: MongoDB provides a powerful query language that allows you to retrive, filter, and manipulate data. You can use the 
    MongoDB Query Language(MQL) to perform operations like find, insert, update and delete.
9.  Aggregation: MongoDB offers an aggregation framework that enables you to perform complex data transformations and calculations on your
    database. It includes operators for grouping, sorting, filtering, and more.
10. Replication: MongoDB supports data replication, allowing you to create multiple copies(replica sets) of your data for high availability
    and fault tolerance. It ensures that data ramains accessible even in the event of server failures.
11. Sharding: For handling large volumes of data, MongoDB offers sharding, a horizontal scaling technique that distributes data across 
    multiple servers or clusters. Sharding helps maintain performance as your dataset grows.
12. Geospatial Queries: MongoDB has built-in support for geospatial data and geospatial indexing, making it suitable for location-based
    applications that require querying based on geographic coordinates.
13. Security: MongoDB provides various securiy features, including authentication, authorization, role-based access control, and encryption
    at rest and in transit, to protect your data.
14. Community and Enterprise Versions: MongoDB is available in both open-source Community Edition and commercial Enterprise Editions, each
    with its own set of features and support options.
15. Ecosystem: MongoDB has a rich ecosystem with drivers and libraries availabe for multiple programming languages, as well as tools for
    import/exoprt, monitoring and management.

* To Perform CRUD in MongoDB
        -------------
# Assume you have collection named users in database mydb.
* To switch database use 'use mydb' . If there is no database named mydb, then a database 'users' will be created with 'use'.

1. Create Data:
   # Use the insertOne() or insertMany() methods
       db.users.insertOne({
            name: "John Doe",
            age: 25,
            email: "johndoe@email.com"
       });

      # insertMany()
       db.users.insertMany([
            {name: "John Doe", age: 30},
            {name: "Doe John", age: 55},
            {name: "Jain Doe", age: 52}
       ]);

2. Read Data:
   # To read all data in a collection
   db.users.find({});
   # To read specific data
   db.users.find({email:"johndoe@email.com"})

3. Update Data:
   # Use the updateOne() or UpdateMany() methods.
   db.users.updateOne(
      { name: "John Doe" },
      { $set: {email: "john@email.com"}}
   );

      # updateMany()
      db.users.updateMany(
            {age: {$gt 30}},             # Query to find users older than 30
            {$set: {status: "senior"}}   # Update to set status to "senior"
      );


4. # Use the deleteOne() or deleteMany() methods.
   db.users.deleteOne({name: "John Doe"});

   # deleteMany()
   db.users.deleteMany({status: "inactive"});   # Delete all documents where the status is "inactive".