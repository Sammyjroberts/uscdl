/**
* ADCSSensorData
* Raw sensor data from ADCS sensors
*/
export interface ADCSSensorData {
  /** Raw magnetometer readings (nT) */
  magnetometerReadings: number[];
  /** Raw sun sensor readings (counts) */
  sunSensorReadings: number[];
  /** Gyroscope readings (rad/s) */
  gyroscopeReadings: number[];
  /** Timestamp of the sensor readings (ms) */
  sensorTimestamp: number;
  /** Bitmask of currently enabled sensors */
  sensorsEnabled: number;
}

/**
* Creates a default ADCSSensorData object
* @returns A new ADCSSensorData with default values
*/
export function createADCSSensorData(): ADCSSensorData {
  return {
    magnetometerReadings: Array(3).fill(0),
    sunSensorReadings: Array(6).fill(0),
    gyroscopeReadings: Array(3).fill(0),
    sensorTimestamp: 0,
    sensorsEnabled: 0
  };
}

/**
* Serializes a ADCSSensorData object to an ArrayBuffer
* @param data The ADCSSensorData object to serialize
* @returns An ArrayBuffer containing the serialized data
*/
export function serializeADCSSensorData(data: ADCSSensorData): ArrayBuffer {
  const buffer = new ArrayBuffer(35);
  const view = new DataView(buffer);
  let offset = 0;
  // Serialize magnetometerReadings array
  for (let i = 0; i < 3; i++) {
    view.setInt16(offset, data.magnetometerReadings[i], true);
    offset += 2;
  }
  // Serialize sunSensorReadings array
  for (let i = 0; i < 6; i++) {
    view.setUint16(offset, data.sunSensorReadings[i], true);
    offset += 2;
  }
  // Serialize gyroscopeReadings array
  for (let i = 0; i < 3; i++) {
    view.setFloat32(offset, data.gyroscopeReadings[i], true);
    offset += 4;
  }
  // Serialize sensorTimestamp scalar
  view.setUint32(offset, data.sensorTimestamp, true);
  offset += 4;
  // Serialize sensorsEnabled scalar
  view.setUint8(offset, data.sensorsEnabled);
  offset += 1;

  return buffer;
}

/**
* Deserializes an ArrayBuffer to a ADCSSensorData object
* @param buffer The ArrayBuffer containing serialized data
* @returns A ADCSSensorData object with the deserialized data
*/
export function deserializeADCSSensorData(buffer: ArrayBuffer): ADCSSensorData {
  const view = new DataView(buffer);
  let offset = 0;
  const result = createADCSSensorData();
  // Deserialize magnetometerReadings array
  const magnetometerReadingsArray = [];
  for (let i = 0; i < 3; i++) {
    magnetometerReadingsArray.push(view.getInt16(offset, true));
    offset += 2;
  }
  result.magnetometerReadings = magnetometerReadingsArray;
  // Deserialize sunSensorReadings array
  const sunSensorReadingsArray = [];
  for (let i = 0; i < 6; i++) {
    sunSensorReadingsArray.push(view.getUint16(offset, true));
    offset += 2;
  }
  result.sunSensorReadings = sunSensorReadingsArray;
  // Deserialize gyroscopeReadings array
  const gyroscopeReadingsArray = [];
  for (let i = 0; i < 3; i++) {
    gyroscopeReadingsArray.push(view.getFloat32(offset, true));
    offset += 4;
  }
  result.gyroscopeReadings = gyroscopeReadingsArray;
  // Deserialize sensorTimestamp scalar
  result.sensorTimestamp = view.getUint32(offset, true);
  offset += 4;
  // Deserialize sensorsEnabled scalar
  result.sensorsEnabled = view.getUint8(offset);
  offset += 1;

  return result;
}
