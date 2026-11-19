# TODO Application — User Stories(new)

## US-01: Create a TODO

**User Story**

As a user, I want to create a TODO task so that I can keep track of work that needs to be completed.

**Acceptance Criteria**

1. As  user I can  create a TODO by providing valid task details.
2. The system assigns a unique ID to each TODO.
3. The system must informs the user  when the TODO is created successfully.
4. system informs the user when the TODO cannot be created because the provided details are invalid.

## US-02: Retrieve All TODO Tasks

**User Story**

As a user, I want to view all my TODO tasks so that I can manage my work in one place.

**Acceptance Criteria**

1. As user i can view the available TODO tasks.
2. Each TODO displays its unique ID, name, and status.
3. If no TODO tasks exist, the system displays an empty list.
4. A user can view tasks page by page when pagination is supported.

## US-03: Retrieve a TODO Task by ID

**User Story**

As a user, I want to view a TODO task by its unique ID so that I can see its details.

**Acceptance Criteria**

1. A user can retrieve a TODO by providing its ID.
2. The system displays the details of the requested TODO.
3. If the TODO does not exist, the system informs the user.
4. If the provided ID is invalid, the system informs the user.

## US-04: Update a TODO Task

**User Story**

As a user, I want to update an existing TODO task so that I can keep its details up to date.

**Acceptance Criteria**

1. A user can update the details of an existing TODO.
2. The system saves the updated details successfully.
3. The TODO retains its unique ID after an update.
4. If the TODO does not exist, the system informs the user.
5. If the provided details are invalid, the system informs the user.

## US-05: Delete a TODO Task

**User Story**

As a user, I want to delete a TODO task so that I can remove tasks that I no longer need.

**Acceptance Criteria**

1. A user can delete an existing TODO by its ID.
2. After deletion, the TODO no longer appears in the available task list.
3. The system informs the user when the TODO is deleted successfully.
4. If the TODO does not exist, the system informs the user.

## US-06: Mark a TODO as Completed or Pending

**User Story**

As a user, I want to mark a TODO as completed or pending so that I can track my progress.

**Acceptance Criteria**

1. A user can mark a TODO as completed.
2. A user can mark a TODO as pending.
3. The system displays the updated status of the TODO.
4. If the TODO does not exist, the system informs the user.
5. The system does not accept an unsupported status.

## US-07: Delete All Completed TODOs

**User Story**

As a user, I want to delete all completed TODO tasks so that I can clean up my task list while keeping unfinished tasks.

**Acceptance Criteria**

1. A user can delete all completed TODO tasks in one operation.
2. Only completed TODOs are deleted.
3. Pending TODOs remain available.
4. If there are no completed TODOs, the system handles the request gracefully.
5. The system informs the user of the operation's outcome.


## US-08: Sort TODO Tasks

**User Story**

As a user, i want sort the tasks so i can organize and find tasks easily 

**Acceptance Criteria**

1. As user i can sort the task by  fields like status, creation date
2. As user i can sort tasks in Asc, desc order
3. A system displays tasks in selected order
4. If there is unsupported sorting option the system should handles it 



## US-09: Archive a  TODO Task

**User Story**

As a user, i want archive my task so that i can remove from the current active task list without deleting it

**Acceptance Criteria**

1. As  user i can Archive a todo task
2. the archived task no longer appear in the current active task
3. the archived task and its details remain available in the archive
4. the system should informs when task is archived 



## US-10: view Archived TODO tasks

**User Story**

As a user i want view my archived TODO tasks so that i can access tasks that i have previously archived

**Acceptance Criteria**

1. As user i can view all archived tasks
2. the system should displays info about the archived task
3. is there is no any archived task system should displays empty list 


## US-11: Restore and Archived TODO Task 

**User Story**

As a user, I want to restore an archived TODO task so that I can return it to my active task list

**Acceptance Criteria**

1. As user i can  restore archived todo task
2. after restore the task must have in active task
3. the system should inform to user the task restored succesfully

## US-12: View TODO Audit History

**User Story**

As a user, I want to view the history of changes made to a TODO task so that I can understand what actions have been performed on it.s

**Acceptance Criteria**

1. A user can view the available history of a TODO task.
2. history shows the actions perfomed 
3. the system should shows who performed the particular action on the specific todo task







