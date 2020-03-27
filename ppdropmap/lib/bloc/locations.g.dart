// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'locations.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_Location _$_$_LocationFromJson(Map<String, dynamic> json) {
  return _$_Location(
    json['id'] as int,
    (json['lat'] as num)?.toDouble(),
    (json['lng'] as num)?.toDouble(),
    json['name'] as String,
    json['adress'] as String,
  );
}

Map<String, dynamic> _$_$_LocationToJson(_$_Location instance) =>
    <String, dynamic>{
      'id': instance.id,
      'lat': instance.lat,
      'lng': instance.lng,
      'name': instance.name,
      'adress': instance.adress,
    };

_$_Locations _$_$_LocationsFromJson(Map<String, dynamic> json) {
  return _$_Locations(
    (json['locations'] as List)
        ?.map((e) =>
            e == null ? null : Location.fromJson(e as Map<String, dynamic>))
        ?.toList(),
  );
}

Map<String, dynamic> _$_$_LocationsToJson(_$_Locations instance) =>
    <String, dynamic>{
      'locations': instance.locations,
    };
