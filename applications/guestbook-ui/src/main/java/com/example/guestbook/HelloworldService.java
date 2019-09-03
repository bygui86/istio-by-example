
package com.example.guestbook;

import java.util.HashMap;
import java.util.Map;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.web.client.HttpStatusCodeException;
import org.springframework.web.client.RestTemplate;

public class HelloworldService {

  private final static Log log = LogFactory.getLog(HelloworldService.class);

  private final RestTemplate restTemplate;
  private final String endpoint;

  public HelloworldService(RestTemplate restTemplate, String endpoint) {
    this.restTemplate = restTemplate;
    this.endpoint = endpoint;
  }

  public Map<String, String> greetingFallback() {

    Map<String, String> response = new HashMap<>();
    response.put("greeting", "Unable to connect");
    return response;
  }

  public Map<String, String> greeting(String name) {

    try {
      return restTemplate.getForObject(endpoint + "/" + name, Map.class);
      
    } catch (HttpStatusCodeException e) {
      log.error("Error from Helloworld Service, falling back", e);
      return greetingFallback();
    }
  }
}
