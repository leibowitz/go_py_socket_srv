import socket
import sys
import tempfile
import os

# Create a UDS socket
sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)

filepath = os.path.join(tempfile.gettempdir(), "socket", "sock_srv")
if not os.path.exists(filepath):
    print "Socket does not exist", filepath
    sys.exit(1)

# Connect the socket to the port where the server is listening
server_address = filepath
print >>sys.stderr, 'connecting to %s' % server_address
try:
    sock.connect(server_address)
except socket.error, msg:
    print >>sys.stderr, msg
    sys.exit(1)

try:
    
    # Send data
    message = 'This is the message.  It will be repeated.'
    print >>sys.stderr, 'sending "%s"' % message
    sock.sendall(message)

    #amount_received = 0
    #amount_expected = len(message)
    
    #while amount_received < amount_expected:
    #    data = sock.recv(16)
    #    amount_received += len(data)
    #    print >>sys.stderr, 'received "%s"' % data

finally:
    print >>sys.stderr, 'closing socket'
    #sock.close()
