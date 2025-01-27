import pika
import logging
import time
import contextlib

def fail_on_error(error, message):
    if error:
        logging.critical(f"{message}: {error}")
        raise SystemExit(1)

def callback(ch, method, properties, body):
    logging.info(f"Received a message: {body.decode()}")
    dot_count = body.decode().count('.')
    time.sleep(dot_count)
    logging.info("Done")
    ch.basic_ack(delivery_tag=method.delivery_tag)

def publish_job(body):
    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        channel = connection.channel()
    except Exception as e:
        return fail_on_error(e, "Failed to connect to RabbitMQ")
    
    with contextlib.closing(connection), contextlib.closing(channel):
        try:
            channel.queue_declare(queue='task_queue', durable=True, auto_delete=False, exclusive=False)
        except Exception as e:
            return fail_on_error(e, "Failed to declare a queue")
        
        try:
            channel.basic_publish(
                exchange='',
                routing_key='task_queue',
                body=body,
                properties=pika.BasicProperties(content_type="text/plain", delivery_mode=2)
            )
            logging.info("Message published successfully")
        except Exception as e:
            return fail_on_error(e, "Failed to publish message")

def main():
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
    
    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        channel = connection.channel()
    except Exception as e:
        fail_on_error(e, "Failed to connect to RabbitMQ")
    
    try:
        channel.queue_declare(queue='task_queue', durable=True, auto_delete=False, exclusive=False)
    except Exception as e:
        fail_on_error(e, "Failed to declare a queue")
    
    try:
        channel.basic_qos(prefetch_count=1, prefetch_size=0, global_qos=False)
    except Exception as e:
        fail_on_error(e, "Failed to set QoS")
    
    try:
        channel.basic_consume(queue='task_queue', on_message_callback=callback, auto_ack=False)
        logging.info(" [*] Waiting for messages. To exit press CTRL+C")
        channel.start_consuming()
    except KeyboardInterrupt:
        logging.info("Stopping consumer...")
        channel.stop_consuming()
    except Exception as e:
        fail_on_error(e, "Failed to register a consumer")
    finally:
        connection.close()

if __name__ == "__main__":
    main()
