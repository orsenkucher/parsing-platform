// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named

part of 'locations.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;
Location _$LocationFromJson(Map<String, dynamic> json) {
  return _Location.fromJson(json);
}

class _$LocationTearOff {
  const _$LocationTearOff();

  _Location call(int id, double lat, double lng, String name, String adress) {
    return _Location(
      id,
      lat,
      lng,
      name,
      adress,
    );
  }
}

// ignore: unused_element
const $Location = _$LocationTearOff();

mixin _$Location {
  int get id;
  double get lat;
  double get lng;
  String get name;
  String get adress;

  Map<String, dynamic> toJson();
  $LocationCopyWith<Location> get copyWith;
}

abstract class $LocationCopyWith<$Res> {
  factory $LocationCopyWith(Location value, $Res Function(Location) then) =
      _$LocationCopyWithImpl<$Res>;
  $Res call({int id, double lat, double lng, String name, String adress});
}

class _$LocationCopyWithImpl<$Res> implements $LocationCopyWith<$Res> {
  _$LocationCopyWithImpl(this._value, this._then);

  final Location _value;
  // ignore: unused_field
  final $Res Function(Location) _then;

  @override
  $Res call({
    Object id = freezed,
    Object lat = freezed,
    Object lng = freezed,
    Object name = freezed,
    Object adress = freezed,
  }) {
    return _then(_value.copyWith(
      id: id == freezed ? _value.id : id as int,
      lat: lat == freezed ? _value.lat : lat as double,
      lng: lng == freezed ? _value.lng : lng as double,
      name: name == freezed ? _value.name : name as String,
      adress: adress == freezed ? _value.adress : adress as String,
    ));
  }
}

abstract class _$LocationCopyWith<$Res> implements $LocationCopyWith<$Res> {
  factory _$LocationCopyWith(_Location value, $Res Function(_Location) then) =
      __$LocationCopyWithImpl<$Res>;
  @override
  $Res call({int id, double lat, double lng, String name, String adress});
}

class __$LocationCopyWithImpl<$Res> extends _$LocationCopyWithImpl<$Res>
    implements _$LocationCopyWith<$Res> {
  __$LocationCopyWithImpl(_Location _value, $Res Function(_Location) _then)
      : super(_value, (v) => _then(v as _Location));

  @override
  _Location get _value => super._value as _Location;

  @override
  $Res call({
    Object id = freezed,
    Object lat = freezed,
    Object lng = freezed,
    Object name = freezed,
    Object adress = freezed,
  }) {
    return _then(_Location(
      id == freezed ? _value.id : id as int,
      lat == freezed ? _value.lat : lat as double,
      lng == freezed ? _value.lng : lng as double,
      name == freezed ? _value.name : name as String,
      adress == freezed ? _value.adress : adress as String,
    ));
  }
}

@JsonSerializable()
class _$_Location with DiagnosticableTreeMixin implements _Location {
  const _$_Location(this.id, this.lat, this.lng, this.name, this.adress)
      : assert(id != null),
        assert(lat != null),
        assert(lng != null),
        assert(name != null),
        assert(adress != null);

  factory _$_Location.fromJson(Map<String, dynamic> json) =>
      _$_$_LocationFromJson(json);

  @override
  final int id;
  @override
  final double lat;
  @override
  final double lng;
  @override
  final String name;
  @override
  final String adress;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'Location(id: $id, lat: $lat, lng: $lng, name: $name, adress: $adress)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'Location'))
      ..add(DiagnosticsProperty('id', id))
      ..add(DiagnosticsProperty('lat', lat))
      ..add(DiagnosticsProperty('lng', lng))
      ..add(DiagnosticsProperty('name', name))
      ..add(DiagnosticsProperty('adress', adress));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _Location &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.lat, lat) ||
                const DeepCollectionEquality().equals(other.lat, lat)) &&
            (identical(other.lng, lng) ||
                const DeepCollectionEquality().equals(other.lng, lng)) &&
            (identical(other.name, name) ||
                const DeepCollectionEquality().equals(other.name, name)) &&
            (identical(other.adress, adress) ||
                const DeepCollectionEquality().equals(other.adress, adress)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(lat) ^
      const DeepCollectionEquality().hash(lng) ^
      const DeepCollectionEquality().hash(name) ^
      const DeepCollectionEquality().hash(adress);

  @override
  _$LocationCopyWith<_Location> get copyWith =>
      __$LocationCopyWithImpl<_Location>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_LocationToJson(this);
  }
}

abstract class _Location implements Location {
  const factory _Location(
      int id, double lat, double lng, String name, String adress) = _$_Location;

  factory _Location.fromJson(Map<String, dynamic> json) = _$_Location.fromJson;

  @override
  int get id;
  @override
  double get lat;
  @override
  double get lng;
  @override
  String get name;
  @override
  String get adress;
  @override
  _$LocationCopyWith<_Location> get copyWith;
}

Locations _$LocationsFromJson(Map<String, dynamic> json) {
  return _Locations.fromJson(json);
}

class _$LocationsTearOff {
  const _$LocationsTearOff();

  _Locations call(List<Location> locations) {
    return _Locations(
      locations,
    );
  }
}

// ignore: unused_element
const $Locations = _$LocationsTearOff();

mixin _$Locations {
  List<Location> get locations;

  Map<String, dynamic> toJson();
  $LocationsCopyWith<Locations> get copyWith;
}

abstract class $LocationsCopyWith<$Res> {
  factory $LocationsCopyWith(Locations value, $Res Function(Locations) then) =
      _$LocationsCopyWithImpl<$Res>;
  $Res call({List<Location> locations});
}

class _$LocationsCopyWithImpl<$Res> implements $LocationsCopyWith<$Res> {
  _$LocationsCopyWithImpl(this._value, this._then);

  final Locations _value;
  // ignore: unused_field
  final $Res Function(Locations) _then;

  @override
  $Res call({
    Object locations = freezed,
  }) {
    return _then(_value.copyWith(
      locations:
          locations == freezed ? _value.locations : locations as List<Location>,
    ));
  }
}

abstract class _$LocationsCopyWith<$Res> implements $LocationsCopyWith<$Res> {
  factory _$LocationsCopyWith(
          _Locations value, $Res Function(_Locations) then) =
      __$LocationsCopyWithImpl<$Res>;
  @override
  $Res call({List<Location> locations});
}

class __$LocationsCopyWithImpl<$Res> extends _$LocationsCopyWithImpl<$Res>
    implements _$LocationsCopyWith<$Res> {
  __$LocationsCopyWithImpl(_Locations _value, $Res Function(_Locations) _then)
      : super(_value, (v) => _then(v as _Locations));

  @override
  _Locations get _value => super._value as _Locations;

  @override
  $Res call({
    Object locations = freezed,
  }) {
    return _then(_Locations(
      locations == freezed ? _value.locations : locations as List<Location>,
    ));
  }
}

@JsonSerializable()
class _$_Locations with DiagnosticableTreeMixin implements _Locations {
  const _$_Locations(this.locations) : assert(locations != null);

  factory _$_Locations.fromJson(Map<String, dynamic> json) =>
      _$_$_LocationsFromJson(json);

  @override
  final List<Location> locations;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'Locations(locations: $locations)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'Locations'))
      ..add(DiagnosticsProperty('locations', locations));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _Locations &&
            (identical(other.locations, locations) ||
                const DeepCollectionEquality()
                    .equals(other.locations, locations)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(locations);

  @override
  _$LocationsCopyWith<_Locations> get copyWith =>
      __$LocationsCopyWithImpl<_Locations>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_LocationsToJson(this);
  }
}

abstract class _Locations implements Locations {
  const factory _Locations(List<Location> locations) = _$_Locations;

  factory _Locations.fromJson(Map<String, dynamic> json) =
      _$_Locations.fromJson;

  @override
  List<Location> get locations;
  @override
  _$LocationsCopyWith<_Locations> get copyWith;
}

class _$LocationsStateTearOff {
  const _$LocationsStateTearOff();

  _SomeLocations some(Locations locations) {
    return _SomeLocations(
      locations,
    );
  }

  _NoneLocations none() {
    return const _NoneLocations();
  }
}

// ignore: unused_element
const $LocationsState = _$LocationsStateTearOff();

mixin _$LocationsState {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result some(Locations locations),
    @required Result none(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result some(Locations locations),
    Result none(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result some(_SomeLocations value),
    @required Result none(_NoneLocations value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result some(_SomeLocations value),
    Result none(_NoneLocations value),
    @required Result orElse(),
  });
}

abstract class $LocationsStateCopyWith<$Res> {
  factory $LocationsStateCopyWith(
          LocationsState value, $Res Function(LocationsState) then) =
      _$LocationsStateCopyWithImpl<$Res>;
}

class _$LocationsStateCopyWithImpl<$Res>
    implements $LocationsStateCopyWith<$Res> {
  _$LocationsStateCopyWithImpl(this._value, this._then);

  final LocationsState _value;
  // ignore: unused_field
  final $Res Function(LocationsState) _then;
}

abstract class _$SomeLocationsCopyWith<$Res> {
  factory _$SomeLocationsCopyWith(
          _SomeLocations value, $Res Function(_SomeLocations) then) =
      __$SomeLocationsCopyWithImpl<$Res>;
  $Res call({Locations locations});

  $LocationsCopyWith<$Res> get locations;
}

class __$SomeLocationsCopyWithImpl<$Res>
    extends _$LocationsStateCopyWithImpl<$Res>
    implements _$SomeLocationsCopyWith<$Res> {
  __$SomeLocationsCopyWithImpl(
      _SomeLocations _value, $Res Function(_SomeLocations) _then)
      : super(_value, (v) => _then(v as _SomeLocations));

  @override
  _SomeLocations get _value => super._value as _SomeLocations;

  @override
  $Res call({
    Object locations = freezed,
  }) {
    return _then(_SomeLocations(
      locations == freezed ? _value.locations : locations as Locations,
    ));
  }

  @override
  $LocationsCopyWith<$Res> get locations {
    if (_value.locations == null) {
      return null;
    }
    return $LocationsCopyWith<$Res>(_value.locations, (value) {
      return _then(_value.copyWith(locations: value));
    });
  }
}

class _$_SomeLocations with DiagnosticableTreeMixin implements _SomeLocations {
  const _$_SomeLocations(this.locations) : assert(locations != null);

  @override
  final Locations locations;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'LocationsState.some(locations: $locations)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'LocationsState.some'))
      ..add(DiagnosticsProperty('locations', locations));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _SomeLocations &&
            (identical(other.locations, locations) ||
                const DeepCollectionEquality()
                    .equals(other.locations, locations)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(locations);

  @override
  _$SomeLocationsCopyWith<_SomeLocations> get copyWith =>
      __$SomeLocationsCopyWithImpl<_SomeLocations>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result some(Locations locations),
    @required Result none(),
  }) {
    assert(some != null);
    assert(none != null);
    return some(locations);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result some(Locations locations),
    Result none(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (some != null) {
      return some(locations);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result some(_SomeLocations value),
    @required Result none(_NoneLocations value),
  }) {
    assert(some != null);
    assert(none != null);
    return some(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result some(_SomeLocations value),
    Result none(_NoneLocations value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (some != null) {
      return some(this);
    }
    return orElse();
  }
}

abstract class _SomeLocations implements LocationsState {
  const factory _SomeLocations(Locations locations) = _$_SomeLocations;

  Locations get locations;
  _$SomeLocationsCopyWith<_SomeLocations> get copyWith;
}

abstract class _$NoneLocationsCopyWith<$Res> {
  factory _$NoneLocationsCopyWith(
          _NoneLocations value, $Res Function(_NoneLocations) then) =
      __$NoneLocationsCopyWithImpl<$Res>;
}

class __$NoneLocationsCopyWithImpl<$Res>
    extends _$LocationsStateCopyWithImpl<$Res>
    implements _$NoneLocationsCopyWith<$Res> {
  __$NoneLocationsCopyWithImpl(
      _NoneLocations _value, $Res Function(_NoneLocations) _then)
      : super(_value, (v) => _then(v as _NoneLocations));

  @override
  _NoneLocations get _value => super._value as _NoneLocations;
}

class _$_NoneLocations with DiagnosticableTreeMixin implements _NoneLocations {
  const _$_NoneLocations();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'LocationsState.none()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'LocationsState.none'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _NoneLocations);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result some(Locations locations),
    @required Result none(),
  }) {
    assert(some != null);
    assert(none != null);
    return none();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result some(Locations locations),
    Result none(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (none != null) {
      return none();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result some(_SomeLocations value),
    @required Result none(_NoneLocations value),
  }) {
    assert(some != null);
    assert(none != null);
    return none(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result some(_SomeLocations value),
    Result none(_NoneLocations value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (none != null) {
      return none(this);
    }
    return orElse();
  }
}

abstract class _NoneLocations implements LocationsState {
  const factory _NoneLocations() = _$_NoneLocations;
}

class _$LocationsEventTearOff {
  const _$LocationsEventTearOff();

  _GetLocations getLocations() {
    return const _GetLocations();
  }
}

// ignore: unused_element
const $LocationsEvent = _$LocationsEventTearOff();

mixin _$LocationsEvent {}

abstract class $LocationsEventCopyWith<$Res> {
  factory $LocationsEventCopyWith(
          LocationsEvent value, $Res Function(LocationsEvent) then) =
      _$LocationsEventCopyWithImpl<$Res>;
}

class _$LocationsEventCopyWithImpl<$Res>
    implements $LocationsEventCopyWith<$Res> {
  _$LocationsEventCopyWithImpl(this._value, this._then);

  final LocationsEvent _value;
  // ignore: unused_field
  final $Res Function(LocationsEvent) _then;
}

abstract class _$GetLocationsCopyWith<$Res> {
  factory _$GetLocationsCopyWith(
          _GetLocations value, $Res Function(_GetLocations) then) =
      __$GetLocationsCopyWithImpl<$Res>;
}

class __$GetLocationsCopyWithImpl<$Res>
    extends _$LocationsEventCopyWithImpl<$Res>
    implements _$GetLocationsCopyWith<$Res> {
  __$GetLocationsCopyWithImpl(
      _GetLocations _value, $Res Function(_GetLocations) _then)
      : super(_value, (v) => _then(v as _GetLocations));

  @override
  _GetLocations get _value => super._value as _GetLocations;
}

class _$_GetLocations with DiagnosticableTreeMixin implements _GetLocations {
  const _$_GetLocations();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'LocationsEvent.getLocations()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'LocationsEvent.getLocations'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _GetLocations);
  }

  @override
  int get hashCode => runtimeType.hashCode;
}

abstract class _GetLocations implements LocationsEvent {
  const factory _GetLocations() = _$_GetLocations;
}
