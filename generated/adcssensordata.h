/**
* ADCSSensorData
* Raw sensor data from ADCS sensors
*/

#ifndef ADCSSENSORDATA_H
#define ADCSSENSORDATA_H

#include <stdint.h>
  #include <stdbool.h>

    /**
    * Raw sensor data from ADCS sensors
    */
    typedef struct {
    /* Raw magnetometer readings (nT) */
    int16_t magnetometerReadings[3];
    /* Raw sun sensor readings (counts) */
    uint16_t sunSensorReadings[6];
    /* Gyroscope readings (rad/s) */
    float gyroscopeReadings[3];
    /* Timestamp of the sensor readings (ms) */
    uint32_t sensorTimestamp;
    /* Bitmask of currently enabled sensors */
    uint8_t sensorsEnabled;
    } ADCSSensorData_t;

    /**
    * Initialize a ADCSSensorData structure with default values
    * @param p_data Pointer to the structure to initialize
    */
    void adcs_sensor_data_init(ADCSSensorData_t* p_data);

    #endif /* ADCSSENSORDATA_H */
    