API http://109.172.89.158:8080/  
  - /plants GET, POST  
  - /plants/{id} GET, PUT, DELETE
  - /metrics (кастомные api_requests_total, api_request_duration_seconds)

Grafana http://109.172.89.158:3000/ (admin, admin)  

# Метрики  
В middleware/metrics.go случайная задержка обработки запроса до 1,5 с.  

![image](https://github.com/user-attachments/assets/3f1ebdbd-1c0f-40db-bab7-93e975059c9b)  

![image](https://github.com/user-attachments/assets/5b30e967-5798-4aa9-9bf9-b8fecb2c4437)  

# Логи  
По /plants симулируются 2 разные ошибки и заносятся в лог  

![image](https://github.com/user-attachments/assets/e8fda427-ead5-4379-8ebe-55d03776b014)  

![image](https://github.com/user-attachments/assets/93dd73eb-3e32-47d0-aa69-6c1c1093ee49)  

![image](https://github.com/user-attachments/assets/676da56c-59a1-47ea-a8c7-71c623ebca2c)  

# Трейсы  
В middleware/traces.go симуляция дополнительной работы последовательно и параллельно с обработкой запроса  

![image](https://github.com/user-attachments/assets/bf6e0f9d-faba-41ad-80a9-a2937f4bc7b3)  

![image](https://github.com/user-attachments/assets/f9d82c43-ab4f-45f2-891f-6cec2a9a38cd)  

![image](https://github.com/user-attachments/assets/910823d4-e710-4e93-b2d4-edc43dcd85f7)



