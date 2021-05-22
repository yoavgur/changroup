# Changroup
Changroup is a minimal implemenation of sync.WaitGroup that allows you to call `Wait()` in a non-blocking manner, by exposing it through a channel.
The purpose of this package is simplify situations where `Wait()` should be performed with a timeout, or when other operations should be carried out
simultaneously (through the use of select).
