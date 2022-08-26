# Comments:
 1. ++ Tests are OK
 2. + In tests we usually use assert. Don't see the reason to use skip.
 3. - All methods A B C must implement the same interface; and return the same data types
 4. ++ Very good approach to use A1 A2 A3 A4
 5. -- We need a single CALL method, which can correctly work with various A implementations.
 I have written a sample , where the implementation is chosen by the strategy parameter.
 
 # So the criteria are still the same
 
 GIVEN four different implementation of A
     1. return "A", nil
     2. return "", error
     3. execute 100/0 inside
     4. time.Sleep(10*time.Second) when function maximum execution time = 2 sec
 
 WHEN
     execute the CALL method with the strategy parameter
     AND
     with call timeout = 2 sec
 THEN
     have the expected results in unit tests