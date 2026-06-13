package sensors

import (
	"time"
	"github.com/Allthebester/byd-hass/internal/location"
)

// SensorData struct to hold all possible sensor values.
// We use pointers to float64 for numeric values so we can distinguish between a missing value (nil) and a value of 0.
// Status-based fields have been converted to *string to allow English labels in Home Assistant.
type SensorData struct {
	Timestamp time.Time `json:"timestamp"`

	// --- Core Vehicle Data ---
	Speed            *float64 `json:"speed,omitempty"`
	Mileage          *float64 `json:"mileage,omitempty"`
	GearPosition     *float64 `json:"gear_position,omitempty"`
	PowerStatus      *string  `json:"power_status,omitempty"` // CONVERTED: On/Off
	SteeringAngle    *float64 `json:"steering_angle,omitempty"`
	AcceleratorDepth *float64 `json:"accelerator_depth,omitempty"`
	BrakeDepth       *float64 `json:"brake_depth,omitempty"`

	// --- Powertrain & Battery ---
	EnginePower           *float64 `json:"engine_power,omitempty"`
	EngineRPM             *float64 `json:"engine_rpm,omitempty"`
	FrontMotorRPM         *float64 `json:"front_motor_rpm,omitempty"`
	FrontMotorTorque      *float64 `json:"front_motor_torque,omitempty"`
	RearMotorRPM          *float64 `json:"rear_motor_rpm,omitempty"`
	FuelPercentage        *float64 `json:"fuel_percentage,omitempty"`
	BatteryPercentage     *float64 `json:"battery_percentage,omitempty"`
	BatteryCapacity       *float64 `json:"battery_capacity,omitempty"`
	ChargingStatus        *string  `json:"charging_status,omitempty"` // CONVERTED: Status labels
	ChargeGunState        *string  `json:"charge_gun_state,omitempty"` // CONVERTED: Connected/Disconnected
	MaxBatteryVoltage     *float64 `json:"max_battery_voltage,omitempty"`
	MinBatteryVoltage     *float64 `json:"min_battery_voltage,omitempty"`
	TotalPowerConsumption *float64 `json:"total_power_consumption,omitempty"`
	PowerConsumption100km *float64 `json:"power_consumption_100km,omitempty"`
	BatteryVoltage12V     *float64 `json:"battery_voltage_12v,omitempty"`

	// --- Temperature Sensors ---
	AvgBatteryTemp     *float64 `json:"avg_battery_temp,omitempty"`
	MinBatteryTemp     *float64 `json:"min_battery_temp,omitempty"`
	MaxBatteryTemp     *float64 `json:"max_battery_temp,omitempty"`
	CabinTemperature   *float64 `json:"cabin_temperature,omitempty"`
	OutsideTemperature *float64 `json:"outside_temperature,omitempty"`
	TemperatureUnit    *string  `json:"temperature_unit,omitempty"` // CONVERTED: °C/°F

	// --- Doors & Locks ---
	DriverDoor         *string `json:"driver_door,omitempty"` // CONVERTED: Open/Closed
	PassengerDoor      *string `json:"passenger_door,omitempty"` // CONVERTED: Open/Closed
	LeftRearDoor       *string `json:"left_rear_door,omitempty"` // CONVERTED: Open/Closed
	RightRearDoor      *string `json:"right_rear_door,omitempty"` // CONVERTED: Open/Closed
	TrunkDoor          *string `json:"trunk_door,omitempty"` // CONVERTED: Open/Closed
	Hood               *string `json:"hood,omitempty"` // CONVERTED: Open/Closed
	DriverDoorLock     *string `json:"driver_door_lock,omitempty"` // CONVERTED: Locked/Unlocked
	PassengerDoorLock  *string `json:"passenger_door_lock,omitempty"` // CONVERTED: Locked/Unlocked
	LeftRearDoorLock   *string `json:"left_rear_door_lock,omitempty"` // CONVERTED: Locked/Unlocked
	RightRearDoorLock  *string `json:"right_rear_door_lock,omitempty"` // CONVERTED: Locked/Unlocked
	TrunkLock          *string `json:"trunk_lock,omitempty"` // CONVERTED: Locked/Unlocked
	RemoteLockStatus   *string `json:"remote_lock_status,omitempty"` // CONVERTED: Locked/Unlocked
	LeftRearChildLock  *string `json:"left_rear_child_lock,omitempty"` // CONVERTED: Locked/Unlocked
	RightRearChildLock *string `json:"right_rear_child_lock,omitempty"` // CONVERTED: Locked/Unlocked

	// --- Windows & Sunroof ---
	DriverWindowOpenPercent    *float64 `json:"driver_window_open_percent,omitempty"`
	PassengerWindowOpenPercent *float64 `json:"passenger_window_open_percent,omitempty"`
	LeftRearWindowOpenPercent  *float64 `json:"left_rear_window_open_percent,omitempty"`
	RightRearWindowOpenPercent *float64 `json:"right_rear_window_open_percent,omitempty"`
	SunroofOpenPercent         *float64 `json:"sunroof_open_percent,omitempty"`
	SunshadeOpenPercent        *float64 `json:"sunshade_open_percent,omitempty"`

	// --- Tire Pressures ---
	LeftFrontTirePressure  *float64 `json:"left_front_tire_pressure,omitempty"`
	RightFrontTirePressure *float64 `json:"right_front_tire_pressure,omitempty"`
	LeftRearTirePressure   *float64 `json:"left_rear_tire_pressure,omitempty"`
	RightRearTirePressure  *float64 `json:"right_rear_tire_pressure,omitempty"`

	// --- Lights & Wipers ---
	LowBeamLights        *string `json:"low_beam_lights,omitempty"` // CONVERTED: On/Off
	HighBeamLights       *string `json:"high_beam_lights,omitempty"` // CONVERTED: On/Off
	FrontFogLights       *string `json:"front_fog_lights,omitempty"` // CONVERTED: On/Off
	RearFogLights        *string `json:"rear_fog_lights,omitempty"` // CONVERTED: On/Off
	ParkingLights        *string `json:"parking_lights,omitempty"` // CONVERTED: On/Off
	DaytimeRunningLights *string `json:"daytime_running_lights,omitempty"` // CONVERTED: On/Off
	LeftTurnSignal       *string `json:"left_turn_signal,omitempty"` // CONVERTED: On/Off
	RightTurnSignal      *string `json:"right_turn_signal,omitempty"` // CONVERTED: On/Off
	HazardLights         *string `json:"hazard_lights,omitempty"` // CONVERTED: On/Off
	WiperGear            *float64 `json:"wiper_gear,omitempty"`
	FrontWiperSpeed      *float64 `json:"front_wiper_speed,omitempty"`
	LastWiperTime        *float64 `json:"last_wiper_time,omitempty"`

	// --- Climate Control (AC) ---
	ACStatus            *string  `json:"ac_status,omitempty"` // CONVERTED: On/Off
	DriverACTemperature *float64 `json:"driver_ac_temperature,omitempty"`
	FanSpeedLevel       *float64 `json:"fan_speed_level,omitempty"`
	ACCirculationMode   *string  `json:"ac_circulation_mode,omitempty"` // CONVERTED: Recirc/Fresh
	ACBlowingMode       *string  `json:"ac_blowing_mode,omitempty"` // CONVERTED: Face/Feet etc
	Weather             *string  `json:"weather,omitempty"` // CONVERTED: Sunny/Rainy etc
	FootwellLights      *string  `json:"footwell_lights,omitempty"` // CONVERTED: On/Off

	// --- Driving Assistance & Safety ---
	ACCCruiseStatus       *string `json:"acc_cruise_status,omitempty"` // CONVERTED: Enabled/Disabled
	LaneKeepAssistStatus  *string `json:"lane_keep_assist_status,omitempty"` // CONVERTED: Enabled/Disabled
	DriverSeatBeltStatus        *string `json:"driver_seatbelt,omitempty"` // CONVERTED: Buckled/Unbuckled
	PassengerSeatbeltWarn *string `json:"passenger_seatbelt_warn,omitempty"` // CONVERTED: Warning/Ok
	Row2LeftSeatbelt      *string `json:"row2_left_seatbelt,omitempty"` // CONVERTED: Buckled/Unbuckled
	Row2RightSeatbelt     *string `json:"row2_right_seatbelt,omitempty"` // CONVERTED: Buckled/Unbuckled
	Row2CenterSeatbelt    *string `json:"row2_center_seatbelt,omitempty"` // CONVERTED: Buckled/Unbuckled
	DistanceToVehicleAhead    *float64 `json:"distance_to_car_ahead,omitempty"`
	LaneCurvature         *float64 `json:"lane_curvature,omitempty"`
	RightLineDistance     *float64 `json:"right_line_distance,omitempty"`
	LeftLineDistance      *float64 `json:"left_line_distance,omitempty"`
	CruiseSwitch          *string  `json:"cruise_switch,omitempty"` // CONVERTED: On/Off
	AutoParking           *string  `json:"auto_parking,omitempty"` // CONVERTED: Enabled/Disabled

	// --- Radar Sensors ---
	RadarFrontLeft          *float64 `json:"radar_front_left,omitempty"`
	RadarFrontRight         *float64 `json:"radar_front_right,omitempty"`
	RadarRearLeft           *float64 `json:"radar_rear_left,omitempty"`
	RadarRearRight          *float64 `json:"radar_rear_right,omitempty"`
	RadarLeft               *float64 `json:"radar_left,omitempty"`
	RadarFrontMidLeft       *float64 `json:"radar_front_mid_left,omitempty"`
	RadarFrontMidRight      *float64 `json:"radar_front_mid_right,omitempty"`
	RadarRearCenter         *float64 `json:"radar_rear_center,omitempty"`
	RearLeftProximityAlert  *string  `json:"rear_left_proximity_alert,omitempty"` // CONVERTED: Warning/Ok
	RearRightProximityAlert *string  `json:"rear_right_proximity_alert,omitempty"` // CONVERTED: Warning/Ok

	// --- Vehicle & System ---
	VehicleOperatingMode    *float64 `json:"vehicle_operating_mode,omitempty"`
	VehicleRunningMode      *string  `json:"vehicle_running_mode,omitempty"` // CONVERTED: Drive Modes
	SurroundViewStatus      *string  `json:"surround_view_status,omitempty"` // CONVERTED: Displayed/Not
	UIConfigVersion         *float64 `json:"ui_config_version,omitempty"`
	SentryModeStatus        *string  `json:"sentry_mode_status,omitempty"` // CONVERTED: On/Off
	PowerOffRecordingConfig *string  `json:"power_off_recording_config,omitempty"` // CONVERTED: On/Off
	PowerOffSentryAlarm     *string  `json:"power_off_sentry_alarm,omitempty"` // CONVERTED: Alarm/No Alarm
	WiFiStatus              *string  `json:"wifi_status,omitempty"` // CONVERTED: Connected/Disconnected
	BluetoothStatus         *string  `json:"bluetooth_status,omitempty"` // CONVERTED: Connected/Disconnected
	BluetoothSignalStrength *float64 `json:"bluetooth_signal_strength,omitempty"`
	WirelessADBSwitch       *string  `json:"wireless_adb_switch,omitempty"` // CONVERTED: On/Off
	SteeringRotationSpeed   *float64 `json:"steering_rotation_speed,omitempty"`

	// --- AI & Video ---
	AIPersonConfidence     *float64 `json:"ai_person_confidence,omitempty"`
	AIVehicleConfidence    *float64 `json:"ai_vehicle_confidence,omitempty"`
	LastSentryTriggerTime  *float64 `json:"last_sentry_trigger_time,omitempty"`
	LastSentryTriggerImage *string  `json:"last_sentry_trigger_image,omitempty"`
	LastVideoStartTime     *float64 `json:"last_video_start_time,omitempty"`
	LastVideoEndTime       *float64 `json:"last_video_end_time,omitempty"`
	LastVideoPath          *string  `json:"last_video_path,omitempty"`

	// --- Location & Time ---
	Location *location.LocationData `json:"location,omitempty"`
	Year     *float64               `json:"year,omitempty"`
	Month    *float64               `json:"month,omitempty"`
	Day      *float64               `json:"day,omitempty"`
	Hour     *float64               `json:"hour,omitempty"`
	Minute   *float64               `json:"minute,omitempty"`
}

// SensorDefinition provides metadata for a sensor.
type SensorDefinition struct {
	ID                int
	FieldName         string
	ChineseName       string
	EnglishName       string
	Category          string // "sensor" or "binary_sensor"
	DeviceClass       string
	UnitOfMeasurement string
	ScaleFactor       float64
}

// ----------------------------------------------------------------------------
// AllSensors
// ------------
// This table contains one entry for every public field in SensorData that we
// want to surface to higher layers (Diplus polling → MQTT discovery → Home
// Assistant).  Each row provides the metadata needed to build the Diplus query
// template, scale raw values, and publish Home-Assistant discovery messages.
//
//	ID            – Stable numerical identifier (starts at 1, never reused)
//	FieldName     – _Exact_ Go struct field in SensorData (PascalCase)
//	ChineseName   – The precise label Diplus uses in its JSON output
//	EnglishName   – Clear English label for UIs / logs
//	Category      – "sensor" or "binary_sensor" (matches HA platform)
//	DeviceClass   – Optional Home-Assistant device_class (speed, voltage, …)
//	Unit          – Unit of measurement (km/h, °C, %, …) – empty if unit-less
//	ScaleFactor   – Multiply raw value by this to obtain the real value (1 = none)
//
// Whenever you add / remove a field in SensorData **make sure** to update this
// slice accordingly; build failures will warn you if you forget.
// ----------------------------------------------------------------------------
var AllSensors = []SensorDefinition{
	{1, "PowerStatus", "电源状态", "Power Status", "binary_sensor", "power", "", 1},
	{2, "Speed", "车速", "Speed", "sensor", "speed", "km/h", 1},
	{3, "Mileage", "里程", "Mileage", "sensor", "distance", "km", 0.1},
	{4, "GearPosition", "档位", "Gear Position", "sensor", "", "", 1},
	{5, "EngineRPM", "发动机转速", "Engine RPM", "sensor", "", "rpm", 1},
	{6, "BrakeDepth", "刹车深度", "Brake Pedal Depth", "sensor", "", "%", 1},
	{7, "AcceleratorDepth", "加速踏板深度", "Accelerator Pedal Depth", "sensor", "", "%", 1},
	{8, "FrontMotorRPM", "前电机转速", "Front Motor RPM", "sensor", "", "rpm", 1},
	{9, "RearMotorRPM", "后电机转速", "Rear Motor RPM", "sensor", "", "rpm", 1},
	{10, "EnginePower", "发动机功率", "Engine Power", "sensor", "power", "kW", 1},
	{11, "FrontMotorTorque", "前电机扭矩", "Front Motor Torque", "sensor", "", "Nm", 1},
	{12, "ChargeGunState", "充电枪插枪状态", "Charge Gun State", "binary_sensor", "plug", "", 1},
	{13, "PowerConsumption100km", "百公里电耗", "Power consumption per 100 kilometers", "sensor", "", "kWh/100km", 1},
	{14, "MaxBatteryTemp", "最高电池温度", "Maximum Battery Temperature", "sensor", "temperature", "°C", 1},
	{15, "AvgBatteryTemp", "平均电池温度", "Average Battery Temperature", "sensor", "temperature", "°C", 1},
	{16, "MinBatteryTemp", "最低电池温度", "Minimum Battery Temperature", "sensor", "temperature", "°C", 1},
	{17, "MaxBatteryVoltage", "最高电池电压", "Max Battery Voltage", "sensor", "voltage", "V", 1},
	{18, "MinBatteryVoltage", "最低电池电压", "Minimum Battery Voltage", "sensor", "voltage", "V", 1},
	{19, "LastWiperTime", "上次雨刮时间", "Last Wiper Time", "sensor", "", "", 1},
	{20, "Weather", "天气", "Weather", "sensor", "", "", 1},
	{21, "DriverSeatBeltStatus", "主驾驶安全带状态", "Driver's seat belt status", "binary_sensor", "safety", "", 1},
	{22, "RemoteLockStatus", "远程锁车状态", "Remote Lock Status", "binary_sensor", "lock", "", 1},
	// what is ID 23 and 24? not documeneted in the spec.
	{25, "CabinTemperature", "车内温度", "Cabin Temperature", "sensor", "temperature", "°C", 1},
	{26, "OutsideTemperature", "车外温度", "Outside Temperature", "sensor", "temperature", "°C", 1},
	{27, "DriverACTemperature", "主驾驶空调温度", "Driver AC temperature", "sensor", "temperature", "°C", 1},
	{28, "TemperatureUnit", "温度单位", "Temperature unit", "sensor", "", "", 1},
	{29, "BatteryCapacity", "电池容量", "Battery Capacity", "sensor", "energy_storage", "kWh", 1},
	{30, "SteeringAngle", "方向盘转角", "Steering Wheel Angle", "sensor", "", "°", 1},
	{31, "SteeringRotationSpeed", "方向盘转速", "Steering Wheel Speed", "sensor", "", "rpm", 1},
	{32, "TotalPowerConsumption", "总电耗", "Total Power Consumption", "sensor", "energy", "kWh", 1},
	{33, "BatteryPercentage", "电量百分比", "Battery Percentage", "sensor", "battery", "%", 1},
	{34, "FuelPercentage", "油量百分比", "Fuel Percentage", "sensor", "battery", "%", 1},
	{35, "TotalFuelConsumption", "总燃油消耗", "Total Fuel Consumption", "sensor", "", "L", 1},
	{36, "LaneLineCurvature", "车道线曲率", "Lane Line Curvature", "sensor", "", "", 1},
	{37, "RightLaneDistance", "右侧线距离", "Right Lane Distance", "sensor", "", "", 1},
	{38, "LeftLaneDistance", "左侧线距离", "Left Lane Distance", "sensor", "", "", 1},
	{39, "BatteryVoltage12V", "蓄电池电压", "Battery Voltage 12V", "sensor", "", "", 1},
	{40, "RadarLeftFront", "雷达左前", "Radar Left Front", "sensor", "distance", "m", 1},
	{41, "RadarRightFront", "雷达右前", "Radar Right Front", "sensor", "distance", "m", 1},
	{42, "RadarLeftRear", "雷达左后", "Radar Left Rear", "sensor", "distance", "m", 1},
	{43, "RadarRightRear", "雷达右后", "Radar Right Rear", "sensor", "distance", "m", 1},
	{44, "RadarLeft", "雷达左", "Radar Left", "sensor", "distance", "m", 1},
	{45, "RadarFrontLeftCenter", "雷达前左中", "Radar Front Left Center", "sensor", "distance", "m", 1},
	{46, "RadarFrontRightCenter", "雷达前右中", "Radar Front Right Center", "sensor", "distance", "m", 1},
	{47, "RadarCenterRear", "雷达中后", "Radar Center Rear", "sensor", "distance", "m", 1},
	{48, "FrontWiperSpeed", "前雨刮速度", "Front Wiper Speed", "sensor", "", "", 1},
	{49, "WiperGear", "雨刮档位", "Wiper Gear", "sensor", "", "", 1},
	{50, "CruiseSwitch", "巡航开关", "Cruise Switch", "binary_sensor", "", "", 1},
	{51, "DistanceToVehicleAhead", "前车距离", "Distance To The Vehicle Ahead", "sensor", "distance", "m", 1},
	{52, "ChargingStatus", "充电状态", "Charging Status", "sensor", "", "", 1},
	{53, "LeftFrontTirePressure", "左前轮气压", "Left Front Tire Pressure", "sensor", "pressure", "bar", 0.01},
	{54, "RightFrontTirePressure", "右前轮气压", "Right Front Tire Pressure", "sensor", "pressure", "bar", 0.01},
	{55, "LeftRearTirePressure", "左后轮气压", "Left Rear Tire Pressure", "sensor", "pressure", "bar", 0.01},
	{56, "RightRearTirePressure", "右后轮气压", "Right Rear Tire Pressure", "sensor", "pressure", "bar", 0.01},
	{57, "LeftTurnSignal", "左转向灯", "Left Turn Signal", "binary_sensor", "light", "", 1},
	{58, "RightTurnSignal", "右转向灯", "Right Turn Signal", "binary_sensor", "light", "", 1},
	{59, "DriverDoorLock", "主驾车门锁", "Driver Door Lock", "binary_sensor", "lock", "", 1},
	// what is ID 60? not documeneted in the spec.
	{61, "DriverWindowOpenPercent", "主驾车窗打开百分比", "Driver Window Open Percentage", "sensor", "", "%", 1},
	{62, "PassengerWindowOpenPercent", "副驾车窗打开百分比", "Passenger Window Open Percentage", "sensor", "", "%", 1},
	{63, "LeftRearWindowOpenPercent", "左后车窗打开百分比", "Left Rear Window Open Percentage", "sensor", "", "%", 1},
	{64, "RightRearWindowOpenPercent", "右后车窗打开百分比", "Right Rear Window Open Percentage", "sensor", "", "%", 1},
	{65, "SunroofOpenPercent", "天窗打开百分比", "Sunroof Open Percentage", "sensor", "", "%", 1},
	{66, "SunshadeOpenPercent", "遮阳帘打开百分比", "Sunshade Open Percentage", "sensor", "", "%", 1},
	{67, "VehicleOperatingMode", "整车工作模式", "Vehicle Working Mode", "sensor", "", "", 1},
	{68, "VehicleRunningMode", "整车运行模式", "Vehicle Operation Mode", "sensor", "", "", 1},
	{69, "Month", "月", "Month", "sensor", "", "", 1},
	{70, "Day", "日", "Day", "sensor", "", "", 1},
	{71, "Hour", "时", "Hour", "sensor", "", "", 1},
	{72, "Minute", "分", "Minute", "sensor", "", "", 1},
	{73, "PassengerSeatBeltWarn", "副驾安全带警告", "Passenger Seat Belt Warning", "binary_sensor", "safety", "", 1},
	{74, "Row2LeftSeatbelt", "二排左安全带", "Second Row Left Seat Belt", "binary_sensor", "safety", "", 1},
	{75, "Row2RightSeatbelt", "二排右安全带", "Second Row Right Seat Belt", "binary_sensor", "safety", "", 1},
	{76, "Row2CenterSeatbelt", "二排中安全带", "Second Row Center Seat Belt", "binary_sensor", "safety", "", 1},
	{77, "ACStatus", "空调状态", "AC Status", "binary_sensor", "running", "", 1},
	{78, "FanSpeedLevel", "风量档位", "Fan Speed Level", "sensor", "", "", 1},
	{79, "ACCirculationMode", "空调循环方式", "AC Circulation Mode", "sensor", "", "", 1},
	{80, "ACBlowingMode", "空调出风模式", "AC Blowing Mode", "sensor", "", "", 1},
	{81, "DriverDoor", "主驾车门", "Driver Door", "binary_sensor", "door", "", 1},
	{82, "PassengerDoor", "副驾车门", "Passenger Door", "binary_sensor", "door", "", 1},
	{83, "LeftRearDoor", "左后车门", "Left Rear Door", "binary_sensor", "door", "", 1},
	{84, "RightRearDoor", "右后车门", "Right Rear Door", "binary_sensor", "door", "", 1},
	{85, "Hood", "引擎盖", "Hood", "binary_sensor", "opening", "", 1},
	{86, "TrunkDoor", "后备箱门", "Trunk", "binary_sensor", "opening", "", 1},
	{87, "FuelTankCap", "油箱盖", "Fuel Tank Cap", "binary_sensor", "opening", "", 1},
	{88, "AutomaticParking", "自动驻车", "Automatic Parking", "binary_sensor", "running", "", 1},
	{89, "ACCCruiseStatus", "ACC巡航状态", "ACC Cruise Status", "binary_sensor", "moving", "", 1},
	{90, "RearLeftProximityAlert", "左后接近告警", "Left Rear Approach Warning", "binary_sensor", "safety", "", 1},
	{91, "RearRightProximityAlert", "右后接近告警", "Right Rear Approach Warning", "binary_sensor", "safety", "", 1},
	{92, "LaneKeepAssistStatus", "车道保持状态", "Lane Keeping Status", "binary_sensor", "running", "", 1},
	{93, "LeftRearDoorLock", "左后车门锁", "Left Rear Door Lock", "binary_sensor", "lock", "", 1},
	{94, "PassengerDoorLock", "副驾车门锁", "Passenger Door Lock", "binary_sensor", "lock", "", 1},
	{95, "RightRearDoorLock", "右后车门锁", "Right Rear Door Lock", "binary_sensor", "lock", "", 1},
	{96, "TrunkDoorLock", "后备箱门锁", "Trunk Door Lock", "binary_sensor", "lock", "", 1},
	{97, "LeftRearChildLock", "左后儿童锁", "Left Rear Child Lock", "binary_sensor", "lock", "", 1},
	{98, "RightRearChildLock", "右后儿童锁", "Right Rear Child Lock", "binary_sensor", "lock", "", 1},
	{99, "ParkingLights", "小灯", "Parking Lights", "binary_sensor", "light", "", 1},
	{100, "LowBeamLights", "近光灯", "Low Beam", "binary_sensor", "light", "", 1},
	{101, "HighBeamLights", "远光灯", "High Beam", "binary_sensor", "", "", 1},
	// what is ID 102 and 103? not documeneted in the spec.
	{104, "FrontFogLamp", "前雾灯", "Front Fog Lamp", "binary_sensor", "light", "", 1},
	{105, "RearFogLamp", "后雾灯", "Rear Fog Lamp", "binary_sensor", "light", "", 1},
	{106, "FootwellLights", "脚照灯", "Footwell Lights", "binary_sensor", "light", "", 1},
	{107, "DaytimeRunningLights", "日行灯", "Daytime Running Lights", "binary_sensor", "light", "", 1},
	{108, "EngineWaterTemperature", "发动机水温", "Engine Water Temperature", "sensor", "", "°C", 1},
	{109, "HazardLights", "双闪", "Hazard Lights", "binary_sensor", "light", "", 1},

	{1001, "SurroundViewStatus", "全景状态", "Panorama Status", "binary_sensor", "running", "", 1},
	{1002, "UIConfigVersion", "配置UI版本", "Configuration UI Version", "sensor", "", "", 1},
	{1003, "SentryModeStatus", "哨兵状态", "Sentry Status", "binary_sensor", "safety", "", 1},
	{1004, "PowerOffRecordingConfig", "熄火录像配置开关", "Recording Configuration Switch", "binary_sensor", "power", "", 1},
	{1006, "PowerOffSentryAlarm", "熄火哨兵报警", "Sentry Alarm", "binary_sensor", "safety", "", 1},
	{1007, "WiFiStatus", "WIFI状态", "WIFI Status", "binary_sensor", "connectivity", "", 1},
	{1008, "BluetoothStatus", "蓝牙状态", "Bluetooth Status", "binary_sensor", "connectivity", "", 1},
	{1009, "BluetoothSignalStrength", "蓝牙信号强度", "Bluetooth Signal Strength", "sensor", "signal_strength", "dBm", 1},
	{1101, "WirelessADBSwitch", "无线ADB开关", "Wireless ADB Switch", "binary_sensor", "switch", "", 1},

	{2001, "AIPersonConfidence", "AI识别人可信度", "AI Person Confidence", "sensor", "", "", 1},
	{2002, "AIVehicleConfidence", "AI识别车可信度", "AI Vehicle Confidence", "sensor", "", "", 1},
	{2003, "LastSentryTriggerTime", "上次哨兵触发时间", "Last Sentry Trigger Time", "sensor", "timestamp", "", 1},
	{2004, "LastSentryTriggerImage", "上次哨兵触发画面", "Last Sentry Trigger Image", "sensor", "", "", 1},
	{2005, "LastVideoStartTime", "上次录像文件开始时间", "Last Video Start Time", "sensor", "timestamp", "", 1},
	{2006, "LastVideoEndTime", "上次录像文件结束时间", "Last Video End Time", "sensor", "timestamp", "", 1},
	{2007, "LastVideoPath", "上次录像路径", "Last Video Path", "sensor", "", "", 1},
}

// GetSensorByID returns a sensor definition by its ID
func GetSensorByID(id int) *SensorDefinition {
	for _, sensor := range AllSensors {
		if sensor.ID == id {
			return &sensor
		}
	}
	return nil
}

// GetScaleFactor returns the scaling factor for a given JSON field key (snake_case).
// If no explicit factor is defined, 1.0 is returned.
func GetScaleFactor(jsonKey string) float64 {
	factor := 1.0
	for _, s := range AllSensors {
		if ToSnakeCase(s.FieldName) == jsonKey {
			if s.ScaleFactor != 0 {
				factor = s.ScaleFactor
			}
		}
	}
	return factor
}
