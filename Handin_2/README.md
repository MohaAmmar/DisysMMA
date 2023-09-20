# Handin2 readme

# a: What are packages in your implementation? What data structure do you use to transmit data and meta-data?
- The packages in our implementation consists of the struct data, which includes 3 fields, seq, ack & msg. seq & ack being integers that are passed to create a connection between client & server and msg that is being passed from client to server after a connection has been established. The data is transmitted as strings.

# b: Does your implementation use threads or processes? Why is it not realistic to use threads?
- Our implementation uses threads. Threads are not realistic to use as they do not represent the issues that can occour in a real network. These issues being message -corruption, -loss, -reordering & -duplication. 

# c: In case the network changes the order in which messages are delivered, how would you handle message re-ordering
- In this case, we would use sequence numbering, to ensure that when the messages are collected, that they can be ordered in the way they were intended to be.

# d: In case messages can be delayed or lost, how does your implementation handle message loss?
- In this case, we would make use of the acknowledgements that are sent for each message. If such an acknowledgement is not received, we would resend the message.

# e: Why is the 3-way handshake important?
- It establishes a connection between two parties and ensures that they are both ready to communicate. It also ensures that the connection is reliable and that the parties are ready to send and receive messages. 