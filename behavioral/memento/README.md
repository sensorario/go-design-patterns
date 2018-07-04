# Behavioral » Memento

## Description

The memento pattern provides the ability to restore an object to its previous
state (undo via rollback). It is implemented with three elements:

 - originator
 - caretaker
 - memento

The originator is an object that has an internal state. The caretaker change
the originator's state, but wants to be able to undo the change. The caretaker
first asks the originator for a memento object. Then it does whatever operation
it was going to do. To roll back it returns the memento object to the
originator.


## Implementation

