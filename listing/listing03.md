<h3>Чем отличаются RWMutex от Mutex?</h3>

Ответ:
```text
mutex - одна горутина имеет право на чтение и запись
rwmutex - одна горутина имеет право на запись, но позволяет нескольким горутинам право на чтение
```