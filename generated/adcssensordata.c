/**
* ADCSSensorData
* Raw sensor data from ADCS sensors
*/

#include "adcssensordata.h"
#include <string.h>
#include <stdlib.h>

void adcs_sensor_data_init(ADCSSensorData_t* p_data) {
    if (p_data == NULL) {
        return;
    }
    memset(p_data->magnetometerReadings, 0, sizeof(p_data->magnetometerReadings));
    memset(p_data->sunSensorReadings, 0, sizeof(p_data->sunSensorReadings));
    memset(p_data->gyroscopeReadings, 0.0, sizeof(p_data->gyroscopeReadings));
    p_data->sensorTimestamp = 0;
    p_data->sensorsEnabled = 0;
}

int adcs_sensor_data_serialize(const ADCSSensorData_t* p_data, uint8_t* buffer, size_t buffer_size) {
    if (p_data == NULL || buffer == NULL) {
        return -1;
    }

    // Ensure buffer is large enough
    if (buffer_size < 35) {
        return -1;
    }

    size_t offset = 0;
    uint8_t* ptr = buffer;
    size_t item_size = 0;
    // Direct copy for little-endian or byte types
    item_size = 2 * 3;
    memcpy(ptr + offset, p_data->magnetometerReadings, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    item_size = 2 * 6;
    memcpy(ptr + offset, p_data->sunSensorReadings, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    item_size = 4 * 3;
    memcpy(ptr + offset, p_data->gyroscopeReadings, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    memcpy(ptr + offset, &p_data->sensorTimestamp, 4);
    offset += 4;
    // Direct copy for little-endian or byte types
    memcpy(ptr + offset, &p_data->sensorsEnabled, 1);
    offset += 1;

    return (int)offset;
}

int adcs_sensor_data_deserialize(ADCSSensorData_t* p_data, const uint8_t* buffer, size_t buffer_size) {
    if (p_data == NULL || buffer == NULL) {
        return -1;
    }

    // Initialize the structure
    adcs_sensor_data_init(p_data);

    size_t offset = 0;
    const uint8_t* ptr = buffer;
    size_t item_size = 0;
    // Direct copy for little-endian or byte types
    item_size = 2 * 3;
    if (offset + item_size > buffer_size) {
        return -1;
    }
    memcpy(p_data->magnetometerReadings, ptr + offset, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    item_size = 2 * 6;
    if (offset + item_size > buffer_size) {
        return -1;
    }
    memcpy(p_data->sunSensorReadings, ptr + offset, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    item_size = 4 * 3;
    if (offset + item_size > buffer_size) {
        return -1;
    }
    memcpy(p_data->gyroscopeReadings, ptr + offset, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    if (offset + 4 > buffer_size) {
        return -1;
    }
    memcpy(&p_data->sensorTimestamp, ptr + offset, 4);
    offset += 4;
    // Direct copy for little-endian or byte types
    if (offset + 1 > buffer_size) {
        return -1;
    }
    memcpy(&p_data->sensorsEnabled, ptr + offset, 1);
    offset += 1;

    return (int)offset;
}
